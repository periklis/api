package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/observatorium/api/authentication"
	"github.com/observatorium/api/httperr"
	"github.com/observatorium/api/rules"
	"gopkg.in/yaml.v2"
)

func enforceLabelsInRules(rawRules rules.Rules, tenantLabel string, tenantID string) error {
	for i := range rawRules.Groups {
		for j := range rawRules.Groups[i].Rules {
			switch r := rawRules.Groups[i].Rules[j].(type) {
			case rules.RecordingRule:
				if r.Labels.AdditionalProperties == nil {
					r.Labels.AdditionalProperties = make(map[string]string)
				}

				r.Labels.AdditionalProperties[tenantLabel] = tenantID
				rawRules.Groups[i].Rules[j] = r
			case rules.AlertingRule:
				if r.Labels.AdditionalProperties == nil {
					r.Labels.AdditionalProperties = make(map[string]string)
				}

				r.Labels.AdditionalProperties[tenantLabel] = tenantID
				rawRules.Groups[i].Rules[j] = r
			}
		}
	}

	return nil
}

func unmarshalRules(r io.Reader) (rules.Rules, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return rules.Rules{}, err
	}

	var rawRules rules.Rules
	if err := yaml.Unmarshal(body, &rawRules); err != nil {
		return rules.Rules{}, err
	}

	return rawRules, nil
}

type rulesHandler struct {
	client      rules.ClientInterface
	logger      log.Logger
	tenantLabel string
	source      string
}

func (rh *rulesHandler) get(w http.ResponseWriter, r *http.Request) {
	tenant, ok := authentication.GetTenant(r.Context())
	if !ok {
		httperr.PrometheusAPIError(w, "error finding tenant", http.StatusUnauthorized)
		return
	}

	id, ok := authentication.GetTenantID(r.Context())
	if !ok {
		httperr.PrometheusAPIError(w, "error finding tenant ID", http.StatusUnauthorized)
		return
	}

	resp, err := rh.client.ListRules(r.Context(), rh.source, tenant)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not list rules", "err", err.Error())

		sc := http.StatusInternalServerError
		if resp != nil {
			sc = resp.StatusCode
		}

		httperr.PrometheusAPIError(w, "error listing rules", sc)

		return
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		switch resp.StatusCode {
		case http.StatusNotFound:
			httperr.PrometheusAPIError(w, "no rules found", resp.StatusCode)
		default:
			httperr.PrometheusAPIError(w, "error listing rules", resp.StatusCode)
		}

		return
	}

	rawRules, err := unmarshalRules(resp.Body)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not unmarshal rules", "err", err.Error())
		httperr.PrometheusAPIError(w, "error unmarshaling rules", http.StatusInternalServerError)

		return
	}

	err = enforceLabelsInRules(rawRules, rh.tenantLabel, id)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not enforce labels in rules", "err", err.Error())
		httperr.PrometheusAPIError(w, "failed to process rules", http.StatusInternalServerError)

		return
	}

	body, err := yaml.Marshal(rawRules)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not marshal rules YAML", "err", err.Error())
		httperr.PrometheusAPIError(w, "error marshaling rules YAML", http.StatusInternalServerError)

		return
	}

	if _, err := w.Write(body); err != nil {
		level.Error(rh.logger).Log("msg", "could not write body", "err", err.Error())
		return
	}
}

func (rh *rulesHandler) put(w http.ResponseWriter, r *http.Request) {
	tenant, ok := authentication.GetTenant(r.Context())
	if !ok {
		httperr.PrometheusAPIError(w, "error finding tenant", http.StatusUnauthorized)
	}

	id, ok := authentication.GetTenantID(r.Context())
	if !ok {
		httperr.PrometheusAPIError(w, "error finding tenant ID", http.StatusUnauthorized)
		return
	}

	rawRules, err := unmarshalRules(r.Body)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not unmarshal rules", "err", err.Error())
		httperr.PrometheusAPIError(w, "error unmarshaling rules", http.StatusInternalServerError)

		return
	}

	err = enforceLabelsInRules(rawRules, rh.tenantLabel, id)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not enforce labels in rules", "err", err.Error())
		httperr.PrometheusAPIError(w, "failed to process rules", http.StatusInternalServerError)

		return
	}

	body, err := yaml.Marshal(rawRules)
	if err != nil {
		level.Error(rh.logger).Log("msg", "could not marshal rules YAML", "err", err.Error())
		httperr.PrometheusAPIError(w, "error marshaling rules YAML", http.StatusInternalServerError)

		return
	}

	resp, err := rh.client.SetRulesWithBody(r.Context(), rh.source, tenant, r.Header.Get("Content-type"), bytes.NewReader(body))
	if err != nil {
		sc := http.StatusInternalServerError
		if resp != nil {
			sc = resp.StatusCode
		}

		level.Error(rh.logger).Log("msg", "could not set rules", "err", err.Error())
		httperr.PrometheusAPIError(w, "error creating rules", sc)

		return
	}

	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)

	if _, err := io.Copy(w, resp.Body); err != nil {
		httperr.PrometheusAPIError(w, "error writing rules response", http.StatusInternalServerError)
		return
	}
}

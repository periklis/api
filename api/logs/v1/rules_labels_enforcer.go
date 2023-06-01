package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ghodss/yaml"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/observatorium/api/authentication"
	"github.com/prometheus/prometheus/model/labels"
)

const (
	contentTypeApplicationJSON = "application/json"
	contentTypeApplicationYAML = "application/yaml"
)

var (
	errUnknownTenantKey        = errors.New("Uknown tenant key")
	errUnknownRulesContentType = errors.New("Unknown rules response content type")
)

type alert struct {
	Labels      labels.Labels `json:"labels"`
	Annotations labels.Labels `json:"annotations"`
	State       string        `json:"state"`
	ActiveAt    *time.Time    `json:"activeAt,omitempty"`
	Value       string        `json:"value"`
}

type alertingRule struct {
	Name        string        `json:"name"`
	Query       string        `json:"query"`
	Duration    float64       `json:"duration"`
	Labels      labels.Labels `json:"labels"`
	Annotations labels.Labels `json:"annotations"`
	Alerts      []*alert      `json:"alerts"`
	Health      string        `json:"health"`
	LastError   string        `json:"lastError,omitempty"`
	// Type of an alertingRule is always "alerting".
	Type string `json:"type"`
}

type recordingRule struct {
	Name      string        `json:"name"`
	Query     string        `json:"query"`
	Labels    labels.Labels `json:"labels,omitempty"`
	Health    string        `json:"health"`
	LastError string        `json:"lastError,omitempty"`
	// Type of a recordingRule is always "recording".
	Type string `json:"type"`
}

type ruleGroup struct {
	Name     string  `json:"name"`
	File     string  `json:"file"`
	Rules    []rule  `json:"rules"`
	Interval float64 `json:"interval"`
}

type rule struct {
	*alertingRule
	*recordingRule
}

func (r *rule) Labels() labels.Labels {
	if r.alertingRule != nil {
		return r.alertingRule.Labels
	}
	return r.recordingRule.Labels
}

// MarshalJSON implements the json.Marshaler interface for rule.
func (r *rule) MarshalJSON() ([]byte, error) {
	if r.alertingRule != nil {
		return json.Marshal(r.alertingRule)
	}
	return json.Marshal(r.recordingRule)
}

// UnmarshalJSON implements the json.Unmarshaler interface for rule.
func (r *rule) UnmarshalJSON(b []byte) error {
	var ruleType struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &ruleType); err != nil {
		return err
	}
	switch ruleType.Type {
	case "alerting":
		var alertingr alertingRule
		if err := json.Unmarshal(b, &alertingr); err != nil {
			return err
		}
		r.alertingRule = &alertingr
	case "recording":
		var recordingr recordingRule
		if err := json.Unmarshal(b, &recordingr); err != nil {
			return err
		}
		r.recordingRule = &recordingr
	default:
		return fmt.Errorf("failed to unmarshal rule: unknown type %q", ruleType.Type)
	}

	return nil
}

type rulesData struct {
	RuleGroups []*ruleGroup `json:"groups,omitempty"`
	Alerts     []*alert     `json:"alerts,omitempty"`
}

type prometheusRulesResponse struct {
	Status    string    `json:"status"`
	Data      rulesData `json:"data"`
	Error     string    `json:"error"`
	ErrorType string    `json:"errorType"`
}

type lokiRule struct {
	Alert       string            `json:"alert,omitempty"`
	Record      string            `json:"record,omitempty"`
	Expr        string            `json:"expr,omitempty"`
	For         string            `json:"for,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

type lokiRuleGroup struct {
	Name     string     `json:"name"`
	Interval string     `json:"interval,omitempty"`
	Limit    int        `json:"limit,omitempty"`
	Rules    []lokiRule `json:"rules"`
}

type lokiRulesResponse = map[string][]lokiRuleGroup

func newModifyResponse(logger log.Logger, labelKeys map[string][]string) func(*http.Response) error {
	return func(res *http.Response) error {
		tenant, ok := authentication.GetTenant(res.Request.Context())
		if !ok {
			return errUnknownTenantKey
		}

		keys, ok := labelKeys[tenant]
		if !ok {
			level.Debug(logger).Log("msg", "Skip applying rule label filters", "tenant", tenant)
			return nil
		}

		var (
			matchers    = extractMatchers(res.Request, keys)
			contentType = res.Header.Get("Content-Type")
		)

		matcherStr := fmt.Sprintf("%s", matchers)
		level.Debug(logger).Log("msg", "filtering using matchers", "tenant", tenant, "matchers", matcherStr)

		body, err := io.ReadAll(res.Body)
		if err != nil {
			level.Error(logger).Log("msg", err)
			return err
		}
		res.Body.Close()

		b, err := filterRules(body, contentType, matchers)
		if err != nil {
			level.Error(logger).Log("msg", err)
			return err
		}

		res.Body = io.NopCloser(bytes.NewReader(b))
		res.ContentLength = int64(len(b))

		return nil
	}
}

func extractMatchers(r *http.Request, l []string) map[string]string {
	queryParams := r.URL.Query()
	matchers := map[string]string{}
	for _, name := range l {
		value := queryParams.Get(name)
		if value != "" {
			matchers[name] = value
		}
	}

	return matchers
}

func filterRules(body []byte, contentType string, matchers map[string]string) ([]byte, error) {
	switch contentType {
	case contentTypeApplicationJSON:
		var res prometheusRulesResponse
		err := json.Unmarshal(body, &res)
		if err != nil {
			return nil, err
		}

		return json.Marshal(filterPrometheusResponse(res, matchers))

	case contentTypeApplicationYAML:
		var res lokiRulesResponse
		if err := yaml.Unmarshal(body, &res); err != nil {
			return nil, err
		}

		return yaml.Marshal(filterLokiRules(res, matchers))

	default:
		return nil, errUnknownRulesContentType
	}
}

func filterPrometheusResponse(res prometheusRulesResponse, matchers map[string]string) prometheusRulesResponse {
	if len(matchers) == 0 {
		res.Data = rulesData{}
		return res
	}

	if len(res.Data.RuleGroups) > 0 {
		filtered := filterPrometheusRuleGroups(res.Data.RuleGroups, matchers)
		res.Data = rulesData{RuleGroups: filtered}
	}

	if len(res.Data.Alerts) > 0 {
		filtered := filterPrometheusAlerts(res.Data.Alerts, matchers)
		res.Data = rulesData{Alerts: filtered}
	}

	return res
}

func filterPrometheusRuleGroups(groups []*ruleGroup, matchers map[string]string) []*ruleGroup {
	var filtered []*ruleGroup

	for _, group := range groups {
		var filteredRules []rule

	rules:
		for _, rule := range group.Rules {
			for key, value := range matchers {
				if !rule.Labels().Has(key) || rule.Labels().Get(key) != value {
					continue rules
				}
			}

			filteredRules = append(filteredRules, rule)
		}

		if len(filteredRules) > 0 {
			group.Rules = filteredRules
			filtered = append(filtered, group)
		}
	}

	return filtered
}

func filterPrometheusAlerts(alerts []*alert, matchers map[string]string) []*alert {
	var filtered []*alert

alerts:
	for _, alert := range alerts {
		for key, value := range matchers {
			if !alert.Labels.Has(key) || alert.Labels.Get(key) != value {
				continue alerts
			}
		}

		filtered = append(filtered, alert)
	}

	return filtered
}

func filterLokiRules(res lokiRulesResponse, matchers map[string]string) lokiRulesResponse {
	if len(matchers) == 0 {
		return nil
	}

	filtered := lokiRulesResponse{}

	for name, groups := range res {
		var filteredGroups []lokiRuleGroup

		for _, group := range groups {
			var filteredRules []lokiRule

		rules:
			for _, rule := range group.Rules {
				for key, value := range matchers {
					val, ok := rule.Labels[key]
					if !ok || val != value {
						continue rules
					}
				}

				filteredRules = append(filteredRules, rule)
			}

			if len(filteredRules) > 0 {
				group.Rules = filteredRules
				filteredGroups = append(filteredGroups, group)
			}
		}

		if len(filteredGroups) > 0 {
			filtered[name] = filteredGroups
		}
	}

	return filtered
}

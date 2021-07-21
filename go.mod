module github.com/observatorium/api

go 1.16

require (
	github.com/brancz/kube-rbac-proxy v0.5.0
	github.com/cloudflare/cfssl v1.4.1
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/efficientgo/e2e v0.11.1
	github.com/efficientgo/tools/core v0.0.0-20210731122119-5d4a0645ce9a
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-chi/chi v4.1.0+incompatible
	github.com/go-chi/httprate v0.4.0
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/grafana/loki v1.6.2-0.20210721070658-cb337c3b710f
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.1-0.20191002090509-6af20e3a5340
	github.com/metalmatze/signal v0.0.0-20201002154727-d0c16e42a3cf
	github.com/oklog/run v1.1.0
	github.com/open-policy-agent/opa v0.23.2
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/prometheus-community/prom-label-proxy v0.3.1-0.20210623095334-9d425172d7bb
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.29.0
	github.com/prometheus/prometheus v1.8.2-0.20210621150501-ff58416a0b02
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0
	go.opentelemetry.io/contrib/propagators v0.20.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/trace v0.20.0
	go.uber.org/automaxprocs v1.2.0
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	google.golang.org/genproto v0.0.0-20210604141403-392c879c8b08
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/square/go-jose.v2 v2.4.1 // indirect
	k8s.io/component-base v0.18.0
)

replace k8s.io/client-go => k8s.io/client-go v0.18.0 // Required by prometheus-operator

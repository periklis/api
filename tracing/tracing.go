package tracing

import (
	"fmt"
	"net"

	propjaeger "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.opentelemetry.io/otel/trace"
)

// EndpointType represents the type of the tracing endpoint.
type EndpointType string

const (
	EndpointTypeCollector EndpointType = "collector"
	EndpointTypeAgent     EndpointType = "agent"
)

// InitTracer creates an OTel TracerProvider that exports the traces to a Jaeger agent/collector.
func InitTracer(
	serviceName string,
	endpoint string,
	endpointType EndpointType,
	samplingFraction float64,
) (tp trace.TracerProvider, err error) {
	tp = trace.NewNoopTracerProvider()

	if endpoint == "" {
		return tp, nil
	}

	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse host and port from tracing endpoint: %w", err)
	}

	endpointOption := jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost(host),
		jaeger.WithAgentPort(port),
	)

	if endpointType == EndpointTypeCollector {
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint))
	}

	exp, err := jaeger.NewRawExporter(endpointOption)
	if err != nil {
		return nil, fmt.Errorf("failed to create new raw tracing exporter: %w", err)
	}

	res := resource.NewWithAttributes(
		semconv.ServiceNameKey.String(serviceName),
	)

	s := sdktrace.ParentBased(
		sdktrace.TraceIDRatioBased(samplingFraction),
	)

	tp = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(s),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propjaeger.Jaeger{},
		propagation.Baggage{},
	))

	return tp, nil
}

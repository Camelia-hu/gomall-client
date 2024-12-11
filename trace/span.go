package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func SpanInit(ctx context.Context, traceName string, spanName string) trace.Span {
	tr := otel.Tracer(traceName)
	_, span := tr.Start(ctx, spanName)
	return span
}

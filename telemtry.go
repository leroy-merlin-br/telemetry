package telemetry

import "github.com/opentracing/opentracing-go"

type TracerOption func(opentracing.Tracer) opentracing.Tracer

func InitTracer(options ...TracerOption) opentracing.Tracer {
	var tracer opentracing.Tracer

	for _, option := range options {
		tracer = option(tracer)
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer
}

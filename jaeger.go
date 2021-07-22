package telemetry

import (
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func WithJaeger(name string) TracerOption {
	return func(tracer opentracing.Tracer) opentracing.Tracer {
		if _, ok := os.LookupEnv("JAEGER_ENABLED"); !ok {
			return tracer
		}

		conf, err := config.FromEnv()
		if err != nil {
			log.Fatalln(err)
		}

		conf.ServiceName = name
		conf.Sampler.Type = "const"
		conf.Sampler.Param = 1
		conf.Reporter.LogSpans = true

		tracer, _, err = conf.NewTracer()
		if err != nil {
			log.Fatalln(err)
		}

		return tracer
	}

}

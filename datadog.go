package telemetry

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	missingDataDogEnvironmentVar     = errors.New("missing `DD_ENV` environment vars")
	missingDataDogTraceAgentPortVar  = errors.New("missing `DD_TRACE_AGENT_PORT` environment vars")
	missingDataDogAgentHost          = errors.New("missing `DD_AGENT_HOST` environment vars")
	missingApplicationReleaseVersion = errors.New("missing `DD_VERSION` environment vars")
	missingServiceReleaseVersion     = errors.New("missing `DD_SERVICE` environment vars")
)

type DatadogOptions struct {
	Service, Env, Port, Host, Version string
}

func (o DatadogOptions) validate() error {
	if o.Service == "" {
		return missingServiceReleaseVersion
	}

	if o.Version == "" {
		return missingApplicationReleaseVersion
	}

	if o.Env == "" {
		return missingDataDogEnvironmentVar
	}

	if o.Port == "" {
		return missingDataDogTraceAgentPortVar
	}

	if o.Host == "" {
		return missingDataDogAgentHost
	}

	return nil
}

func WithDatadog(log hclog.Logger, options *DatadogOptions) TracerOption {
	return func(o opentracing.Tracer) opentracing.Tracer {
		if _, ok := os.LookupEnv("DATADOG_ENABLED"); !ok {
			return o
		}

		if err := options.validate(); err != nil {
			log.Error("Tracing options invalid", err)
			panic(err)
		}

		return opentracer.New(
			tracer.WithService(log.Name()),
			tracer.WithServiceVersion(options.Version),
			tracer.WithEnv(options.Env),
			tracer.WithServiceVersion(options.Version),
			tracer.WithAgentAddr(net.JoinHostPort(options.Host, options.Port)),
			tracer.WithSampler(tracer.NewRateSampler(1)),
		)

	}
}

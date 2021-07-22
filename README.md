# telemetry

Opentracing factory to setup by multiples vendors

## Usage
```go
telemetry.InitTracer(
		telemetry.WithJaeger(os.Getenv("APP_NAME")),
		telemetry.WithDatadog(logger, &telemetry.DatadogOptions{
			Service: os.Getenv("APP_NAME"),
			Env:     os.Getenv("DD_ENV"),
			Port:    os.Getenv("DD_TRACE_AGENT_PORT"),
			Host:    os.Getenv("DD_AGENT_HOST"),
			Version: os.Getenv("DD_VERSION"),
		}),
	)
```

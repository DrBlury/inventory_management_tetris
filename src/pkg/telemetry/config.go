package telemetry

type Config struct {
	EnableOTel     bool
	MeterProvider  string
	TracerProvider string
	ServiceName    string
	ServiceVersion string
}

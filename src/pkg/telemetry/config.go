package telemetry

type Config struct {
	DisableOTel    bool
	MeterProvider  string
	TracerProvider string
}

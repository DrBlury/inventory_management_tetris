package telemetry

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

type OtelService struct {
	cfg *Config
	ctx context.Context
}

func NewOtelService(cfg *Config, ctx context.Context) *OtelService {
	otelService := OtelService{
		cfg: cfg,
		ctx: ctx,
	}

	return &otelService
}

// SetupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func (o *OtelService) SetupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// set default to console
	tracerProvider, err := newConsoleTraceProvider()
	if err != nil {
		handleErr(err)
		return
	}
	// Override with none if set
	switch o.cfg.TracerProvider {
	case "console":
		zap.L().Info("OTEL_TRACER_PROVIDER environment variable set to console, using default console exporter")
	case "otlp":
		// OTLP exporter
		// TODO: Add OTLP exporter
		zap.L().Info("OTEL_TRACER_PROVIDER environment variable set to otlp")
	// Add other exporters here
	default:
		zap.L().Info("OTEL_TRACER_PROVIDER environment variable set to unknown value, using default console exporter")
	}

	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	otel.SetTracerProvider(tracerProvider)

	// set default to console
	meterProvider, err := newConsoleMeterProvider()
	if err != nil {
		handleErr(err)
		return
	}
	// Override with none if set
	switch o.cfg.MeterProvider {
	case "console":
		zap.L().Info("OTEL_METER_PROVIDER environment variable set to console, using default console exporter")
	case "otlp":
		// OTLP exporter
		// TODO: Add OTLP exporter
		zap.L().Info("OTEL_METER_PROVIDER environment variable set to otlp")
	// Add other exporters here
	default:
		zap.L().Info("OTEL_METER_PROVIDER environment variable set to unknown value, using default console exporter")
	}

	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)
	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newConsoleTraceProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			trace.WithBatchTimeout(5*time.Second)),
	)
	return traceProvider, nil
}

func newConsoleMeterProvider() (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			metric.WithInterval(5*time.Minute))),
	)
	return meterProvider, nil
}

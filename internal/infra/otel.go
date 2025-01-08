package infra

import (
	"context"
	"encoding/base64"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
	"time"
)

func NewOtel(config *conf.OTELConfig, traceName string) primitive.CloseFunc {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.Username+":"+config.Password))

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(config.Host),
		otlptracegrpc.WithHeaders(map[string]string{
			"Authorization": authHeader,
		}),
	)

	traceExporter, err := otlptrace.New(ctx, traceClient)
	util.Panic(err)

	traceProvider, closeFN, err := startingOtelProvider(ctx, traceExporter, traceName)
	util.Panic(err)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(traceProvider)

	log.Info().Msg("initialization opentelemetry successfully")
	return closeFN
}

func startingOtelProvider(ctx context.Context, exp trace.SpanExporter, traceName string) (*trace.TracerProvider, primitive.CloseFunc, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(traceName),
		),
		resource.WithHost(),
		resource.WithTelemetrySDK(),
		resource.WithFromEnv(),
	)
	if err != nil {
		return nil, nil, err
	}

	bsp := trace.NewBatchSpanProcessor(exp)

	provider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(bsp),
		trace.WithResource(res),
	)
	closeFn := func(ctx context.Context) error {
		log.Info().Msg("starting shutdown export and provider")
		ctxClosure, cancelClosure := context.WithTimeout(ctx, 5*time.Second)
		defer cancelClosure()

		if err = exp.Shutdown(ctxClosure); err != nil {
			return err
		}

		if err = provider.Shutdown(ctxClosure); err != nil {
			return err
		}

		log.Info().Msg("shutdown export and provider successfully")
		return nil
	}

	return provider, closeFn, nil
}

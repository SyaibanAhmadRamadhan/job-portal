package job

import (
	"context"
	"errors"
	"fmt"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/proto/job_post_etl_payload"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/index_jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
	"go.opentelemetry.io/otel/trace"
)

func (s *service) ConsumerPostJobETL(ctx context.Context) (err error) {
	consumer, err := s.eventBusConsumerRepository.ConsumerJobPostETL(ctx)
	if err != nil {
		return tracer.ErrInternalServer(err)
	}

	propagator := otel.GetTextMapPropagator()

	for {
		var data job_post_etl_payload.JobPost
		msg, err := consumer.Reader.FetchMessage(ctx, &data)
		if err != nil {
			if !errors.Is(err, ekafka.ErrJsonUnmarshal) {
				return tracer.ErrInternalServer(err)
			}
			continue
		}

		carrier := ekafka.NewMsgCarrier(&msg)
		ctxConsumer := propagator.Extract(context.Background(), carrier)

		ctxConsumer, span := otel.Tracer("").Start(ctxConsumer, fmt.Sprintf("process cdc job post data. %s - %s", data.Id, data.Title),
			trace.WithAttributes(
				attribute.String("cdc.data.id", data.Id),
				attribute.String("cdc.data.title", data.Title),
				attribute.String("cdc.data.company.id", data.Company.Id),
			))
		span.End()

		err = s.indexJobRepository.CreateNewRecord(ctxConsumer, index_jobs.CreateNewRecordInput{
			Data: entity.IndexJobEntity{
				Company: entity.IndexJobCompany{
					ID:   data.Company.Id,
					Name: data.Company.Name,
				},
				ID:          data.Id,
				Title:       data.Title,
				Description: data.Description,
				Timestamp:   data.CreatedAt.AsTime(),
			},
		})
		if err != nil {
			span.RecordError(tracer.ErrInternalServer(err))
			span.SetStatus(codes.Error, err.Error())
			span.SetAttributes(semconv.ErrorTypeKey.String("failed create job post etl"))
			span.End()
			return tracer.ErrInternalServer(err)
		}

		err = consumer.Reader.CommitMessages(ctx, msg)
		if err != nil {
			span.RecordError(tracer.ErrInternalServer(err))
			span.SetStatus(codes.Error, err.Error())
			span.SetAttributes(semconv.ErrorTypeKey.String("failed commit message"))
			span.End()
			return tracer.ErrInternalServer(err)
		}
		span.SetStatus(codes.Ok, "cdc successfully")
		span.End()
	}
}

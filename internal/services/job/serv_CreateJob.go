package job

import (
	"context"
	"database/sql"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/proto/job_post_etl_payload"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *service) CreateJob(ctx context.Context, input CreateJobInput) (output CreateJobOutput, err error) {
	err = s.dbTx.DoTxContext(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false},
		func(ctx context.Context, tx wsqlx.Rdbms) (err error) {
			createCompanyOutput, err := s.companyRepository.CreateNewRecord(ctx, companies.CreateNewRecordInput{
				Tx:   tx,
				Name: input.CompanyName,
			})
			if err != nil {
				return tracer.StackTrace(err)
			}

			createJobOutput, err := s.jobRepository.CreateNewRecord(ctx, jobs.CreateNewRecordInput{
				Tx:          tx,
				CompanyID:   createCompanyOutput.ID,
				Title:       input.Title,
				Description: input.Description,
			})
			if err != nil {
				return tracer.StackTrace(err)
			}

			err = s.companyRepository.ClearCache(ctx)
			if err != nil {
				return tracer.ErrInternalServer(err)
			}

			err = s.eventBusPublisherRepository.PublishJobPostETL(ctx, eventbus.PublishJobPostETLInput{
				Payload: &job_post_etl_payload.JobPost{
					Id: createJobOutput.ID,
					Company: &job_post_etl_payload.Company{
						Id:   createCompanyOutput.ID,
						Name: input.CompanyName,
					},
					Title:       input.Title,
					Description: input.Description,
					CreatedAt:   timestamppb.New(createJobOutput.Timestamp),
				},
			})
			if err != nil {
				return tracer.StackTrace(err)
			}

			output = CreateJobOutput{
				CompanyID: createCompanyOutput.ID,
				JobID:     createJobOutput.ID,
			}
			return
		},
	)
	if err != nil {
		return output, tracer.StackTrace(err)
	}
	return
}

type CreateJobInput struct {
	CompanyName string
	Title       string
	Description string
}

type CreateJobOutput struct {
	CompanyID string
	JobID     string
}

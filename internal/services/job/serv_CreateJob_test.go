package job_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/proto/job_post_etl_payload"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func Test_service_CreateJob(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockCompanyRepository := companies.NewMockCompanyRepository(mock)
	mockJobRepository := jobs.NewMockJobRepository(mock)
	mockEventBusPublisherRepository := eventbus.NewMockPublisherRepository(mock)
	mockDBTxRepository := wsqlx.NewMockTx(mock)
	ctx := context.Background()

	s := job.New(job.Options{
		JobRepository:               mockJobRepository,
		CompanyRepository:           mockCompanyRepository,
		EventBusPublisherRepository: mockEventBusPublisherRepository,
		DBTx:                        mockDBTxRepository,
	})

	t.Run("should be return correct", func(t *testing.T) {
		expectedTx := wsqlx.NewRdbms(nil)

		expectedTimestamp := time.Now().UTC()
		expectedInput := job.CreateJobInput{
			CompanyName: "REDIKRU",
			Title:       "backend developer",
			Description: "minimum experience is 2 years",
		}
		expectedOutput := job.CreateJobOutput{
			CompanyID: uuid.NewString(),
			JobID:     uuid.NewString(),
		}

		mockDBTxRepository.EXPECT().
			DoTxContext(gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, opt *sql.TxOptions, fn func(ctx context.Context, tx wsqlx.Rdbms) error) error {
				mockCompanyRepository.EXPECT().
					CreateNewRecord(ctx, companies.CreateNewRecordInput{
						Tx:   expectedTx,
						Name: expectedInput.CompanyName,
					}).
					Return(companies.CreateNewRecordOutput{
						ID: expectedOutput.CompanyID,
					}, nil)

				mockJobRepository.EXPECT().
					CreateNewRecord(ctx, jobs.CreateNewRecordInput{
						Tx:          expectedTx,
						CompanyID:   expectedOutput.CompanyID,
						Title:       expectedInput.Title,
						Description: expectedInput.Description,
					}).
					Return(jobs.CreateNewRecordOutput{
						ID:        expectedOutput.JobID,
						Timestamp: expectedTimestamp,
					}, nil)

				mockCompanyRepository.EXPECT().
					ClearCache(ctx).Return(nil)

				mockEventBusPublisherRepository.EXPECT().
					PublishJobPostETL(ctx, eventbus.PublishJobPostETLInput{
						Payload: &job_post_etl_payload.JobPost{
							Id: expectedOutput.JobID,
							Company: &job_post_etl_payload.Company{
								Id:   expectedOutput.CompanyID,
								Name: expectedInput.CompanyName,
							},
							Title:       expectedInput.Title,
							Description: expectedInput.Description,
							CreatedAt:   timestamppb.New(expectedTimestamp),
						},
					}).Return(nil)
				return fn(ctx, expectedTx)
			})

		output, err := s.CreateJob(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})

	t.Run("should be return error when create company", func(t *testing.T) {
		expectedTx := wsqlx.NewRdbms(nil)

		expectedInput := job.CreateJobInput{
			CompanyName: "REDIKRU",
			Title:       "backend developer",
			Description: "minimum experience is 2 years",
		}

		mockDBTxRepository.EXPECT().
			DoTxContext(gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(func(ctx context.Context, opt *sql.TxOptions, fn func(ctx context.Context, tx wsqlx.Rdbms) error) error {
				mockCompanyRepository.EXPECT().
					CreateNewRecord(ctx, companies.CreateNewRecordInput{
						Tx:   expectedTx,
						Name: expectedInput.CompanyName,
					}).
					Return(companies.CreateNewRecordOutput{}, errors.New("dummy error"))

				return fn(ctx, expectedTx)
			})

		output, err := s.CreateJob(ctx, expectedInput)
		require.Error(t, err)
		require.Empty(t, output)
		var expectedErr *tracer.ErrTrace
		require.ErrorAs(t, err, &expectedErr)
	})
}

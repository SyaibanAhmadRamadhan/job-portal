package eventbus_test

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/proto/job_post_etl_payload"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func Test_repository_PublishJobPostETL(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockKafkaPubSub := ekafka.NewMockKafkaPubSub(mock)

	ctx := context.TODO()
	repo := eventbus.New(mockKafkaPubSub, &conf.KafkaConfig{
		Host: "localhost",
	})

	t.Run("should be return correct", func(t *testing.T) {
		expectedInput := eventbus.PublishJobPostETLInput{
			Payload: &job_post_etl_payload.JobPost{
				Id: uuid.NewString(),
				Company: &job_post_etl_payload.Company{
					Id:   uuid.NewString(),
					Name: "REDIKRU",
				},
				Title:       "Title",
				Description: "Description",
				CreatedAt:   timestamppb.New(time.Now().UTC()),
			},
		}

		mockKafkaPubSub.EXPECT().
			Publish(ctx, gomock.Any()).Return(ekafka.PubOutput{}, nil)

		err := repo.PublishJobPostETL(ctx, expectedInput)
		require.NoError(t, err)
	})
}

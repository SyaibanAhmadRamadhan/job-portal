package eventbus_test

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_repository_ConsumerJobPostETL(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockKafkaPubSub := ekafka.NewMockKafkaPubSub(mock)

	ctx := context.TODO()
	repo := eventbus.New(mockKafkaPubSub, &conf.KafkaConfig{
		Host:     "localhost",
		Username: "user",
		Password: "pass",
	})

	t.Run("should be return correct", func(t *testing.T) {
		mockKafkaPubSub.EXPECT().
			Subscribe(ctx, gomock.Any()).Return(ekafka.SubOutput{}, nil)

		_, err := repo.ConsumerJobPostETL(ctx)
		require.NoError(t, err)
	})
}

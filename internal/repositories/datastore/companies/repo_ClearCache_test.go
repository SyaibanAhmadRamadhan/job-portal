package companies_test

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/redis/rueidis/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_repository_ClearCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := mock.NewClient(ctrl)

	repo := companies.New(nil, client)

	t.Run("should be return correct", func(t *testing.T) {
		prefix := "companies*"
		mockKeys := []string{"companies:1", "companies:2"}

		client.EXPECT().Do(ctx,
			client.B().Scan().Cursor(0).Match(prefix).Count(100).Build(),
		).Return(mock.Result(mock.RedisArray(
			mock.RedisInt64(0),
			mock.RedisArray(
				mock.RedisString(mockKeys[0]),
				mock.RedisString(mockKeys[1]),
			),
		)))

		client.EXPECT().Do(ctx,
			client.B().Del().Key(mockKeys[0]).Build(),
		).Return(mock.ErrorResult(nil))
		client.EXPECT().Do(ctx,
			client.B().Del().Key(mockKeys[1]).Build(),
		).Return(mock.ErrorResult(nil))

		err := repo.ClearCache(ctx)
		require.NoError(t, err)

	})
}

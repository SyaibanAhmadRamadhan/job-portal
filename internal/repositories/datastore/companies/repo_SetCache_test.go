package companies_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/google/uuid"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func Test_repository_SetCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := mock.NewClient(ctrl)

	repo := companies.New(nil, client)

	t.Run("should be return correct", func(t *testing.T) {
		expectedInput := companies.SetCacheInput{
			ExpiredAt: time.Now().UTC().Add(30 * time.Minute),
			Payload: entity.CompanyCacheEntity{
				Pagination: primitive.PaginationOutput{
					Page:      1,
					PageSize:  10,
					PageCount: 1,
					TotalData: 3,
				},
				Items: []entity.CompanyEntity{
					{
						ID:   uuid.NewString(),
						Name: "KREDITKRU",
					},
					{
						ID:   uuid.NewString(),
						Name: "Goto Group",
					},
					{
						ID:   uuid.NewString(),
						Name: "Shopee",
					},
				},
			},
		}

		key := fmt.Sprintf("companies:%d:%d", expectedInput.Payload.Pagination.Page, expectedInput.Payload.Pagination.PageSize)
		payload, err := json.Marshal(expectedInput.Payload)
		require.NoError(t, err)

		client.EXPECT().Do(ctx,
			client.B().
				Set().
				Key(key).
				Value(rueidis.BinaryString(payload)).
				Exat(expectedInput.ExpiredAt).
				Build(),
		).Return(mock.ErrorResult(nil))

		err = repo.SetCache(ctx, expectedInput)
		require.NoError(t, err)
	})
}

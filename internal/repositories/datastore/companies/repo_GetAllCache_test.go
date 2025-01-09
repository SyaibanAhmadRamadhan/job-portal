package companies_test

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_repository_GetAllCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := mock.NewClient(ctrl)

	repo := companies.New(nil, client)

	t.Run("should be return correct", func(t *testing.T) {
		expectedResultStr := `{"pagination":{"Page":1,"PageSize":10,"PageCount":1,"TotalData":3},"items":[{"id":"fec7c493-845e-446a-83f8-f08e98fbdd88","name":"KREDITKRU"},{"id":"c7a5cdd7-1a67-4e16-92d9-e9411d6ba747","name":"Goto Group"},{"id":"4c52de71-d8fe-452f-9fc5-e7e0b6a2d9ce","name":"Shopee"}]}`
		expectedOutput := companies.GetAllCacheOutput{
			Data: entity.CompanyCacheEntity{
				Pagination: primitive.PaginationOutput{
					Page:      1,
					PageSize:  10,
					PageCount: 1,
					TotalData: 3,
				},
				Items: []entity.CompanyEntity{
					{
						ID:   "fec7c493-845e-446a-83f8-f08e98fbdd88",
						Name: "KREDITKRU",
					},
					{
						ID:   "c7a5cdd7-1a67-4e16-92d9-e9411d6ba747",
						Name: "Goto Group",
					},
					{
						ID:   "4c52de71-d8fe-452f-9fc5-e7e0b6a2d9ce",
						Name: "Shopee",
					},
				},
			},
		}
		expectedInput := companies.GetAllCacheInput{
			Pagination: primitive.PaginationInput{
				Page:     1,
				PageSize: 10,
			},
		}

		client.EXPECT().Do(ctx, client.B().
			Get().Key("companies:1:10").Build()).
			Return(mock.Result(mock.RedisString(expectedResultStr)))

		output, err := repo.GetAllCache(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})

	t.Run("should be return error data is nil", func(t *testing.T) {
		expectedInput := companies.GetAllCacheInput{
			Pagination: primitive.PaginationInput{
				Page:     1,
				PageSize: 10,
			},
		}

		client.EXPECT().Do(ctx, client.B().
			Get().Key("companies:1:10").Build()).
			Return(mock.ErrorResult(rueidis.Nil))

		output, err := repo.GetAllCache(ctx, expectedInput)
		require.Error(t, err)
		require.Empty(t, output)
	})
}

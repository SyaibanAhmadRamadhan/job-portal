package index_jobs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/guregu/null/v5"
	"net/http"
)

func (r *repository) GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error) {
	queryMust := make([]types.Query, 0)
	filterQuery := make([]types.Query, 0)

	if input.SearchKeyword.Valid {
		queryMust = append(queryMust, types.Query{
			Bool: &types.BoolQuery{
				MinimumShouldMatch: 1,
				Should: []types.Query{
					{
						Match: map[string]types.MatchQuery{
							"title": {
								Query: input.SearchKeyword.String,
							},
						},
					},
					{
						Match: map[string]types.MatchQuery{
							"description": {
								Query: input.SearchKeyword.String,
							},
						},
					},
				},
			},
		})
	}

	if input.CompanyID.Valid {
		filterQuery = append(filterQuery, types.Query{
			Term: map[string]types.TermQuery{
				"company.id": {
					Value: input.CompanyID.String,
				},
			},
		})
	}

	queryBool := &types.BoolQuery{}
	if len(queryMust) > 0 {
		queryBool.Must = queryMust
	}
	if len(filterQuery) > 0 {
		queryBool.Filter = filterQuery
	}

	offset := primitive.GetOffsetValue(
		input.Pagination.Page,
		input.Pagination.PageSize,
	)

	searchQuery := &search.Request{
		From: &offset,
		Query: &types.Query{
			Bool: queryBool,
		},
		Size:        &input.Pagination.PageSize,
		TrackScores: null.BoolFrom(true).Ptr(),
		Sort: []types.SortCombinations{map[string]any{
			"_score": map[string]string{
				"order": "desc",
			},
		}},
	}
	res, err := r.client.Search().Index(r.indexName).
		Request(searchQuery).
		Do(ctx)
	if err != nil {
		var errEs *types.ElasticsearchError
		if errors.As(err, &errEs) {
			if errEs.Status == http.StatusBadRequest {
				return output, nil
			}
		}
		return output, tracer.ErrInternalServer(err)
	}

	total := 0
	if res.Hits.Total != nil {
		total = int(res.Hits.Total.Value)
	}

	output = GetAllOutput{
		Pagination: primitive.CreatePaginationOutput(input.Pagination, total),
		Items:      make([]entity.IndexJobEntity, 0),
	}

	for _, hit := range res.Hits.Hits {
		item := entity.IndexJobEntity{}
		err = json.Unmarshal(hit.Source_, &item)
		if err != nil {
			return output, tracer.ErrInternalServer(err)
		}
		output.Items = append(output.Items, item)
	}
	return
}

type GetAllInput struct {
	Pagination    primitive.PaginationInput
	SearchKeyword null.String
	CompanyID     null.String
}

type GetAllOutput struct {
	Pagination primitive.PaginationOutput
	Items      []entity.IndexJobEntity
}

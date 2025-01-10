package index_jobs

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyquery"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"net/http"
)

func (r *repository) CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (err error) {
	_, err = r.client.DeleteByQuery(r.indexName).
		Request(&deletebyquery.Request{
			Query: &types.Query{
				Term: map[string]types.TermQuery{
					"id": {
						Value: input.Data.ID,
					},
				},
			},
		}).
		Do(ctx)
	if err != nil {
		var errEs *types.ElasticsearchError
		if errors.As(err, &errEs) {
			if errEs.Status != http.StatusBadRequest {
				return tracer.ErrInternalServer(err)
			}
		} else {
			return tracer.ErrInternalServer(err)
		}
	}

	_, err = r.client.Index(r.indexName).
		Request(input.Data).
		Do(ctx)
	if err != nil {
		return tracer.ErrInternalServer(err)
	}
	return
}

type CreateNewRecordInput struct {
	Data entity.IndexJobEntity
}

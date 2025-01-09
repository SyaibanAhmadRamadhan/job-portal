package companies

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/redis/rueidis"
	"io"
)

func (r *repository) GetAllCache(ctx context.Context, input GetAllCacheInput) (output GetAllCacheOutput, err error) {
	key := keyFormatter(input.Pagination.Page, input.Pagination.PageSize)
	resp := r.redis.Do(ctx, r.redis.B().
		Get().Key(key).Build(),
	)

	data, err := resp.AsReader()
	if err != nil {
		if errors.Is(err, rueidis.Nil) {
			return output, tracer.ErrNotFound(err)
		}
		return output, tracer.ErrInternalServer(err)
	}

	body, err := io.ReadAll(data)
	err = json.Unmarshal(body, &output.Data)
	if err != nil {
		return output, tracer.ErrInternalServer(err)
	}
	return
}

type GetAllCacheInput struct {
	Pagination primitive.PaginationInput
}

type GetAllCacheOutput struct {
	Data entity.CompanyCacheEntity
}

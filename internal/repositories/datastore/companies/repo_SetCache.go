package companies

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/redis/rueidis"
	"time"
)

func keyFormatter(page int, pageSize int) string {
	return fmt.Sprintf("companies:%d:%d", page, pageSize)
}

func (r *repository) SetCache(ctx context.Context, input SetCacheInput) (err error) {
	payload, err := json.Marshal(input.Payload)
	if err != nil {
		return tracer.ErrInternalServer(err)
	}

	key := keyFormatter(input.Payload.Pagination.Page, input.Payload.Pagination.PageSize)

	err = r.redis.Do(ctx,
		r.redis.B().
			Set().
			Key(key).
			Value(rueidis.BinaryString(payload)).
			Exat(input.ExpiredAt).
			Build(),
	).Error()
	return
}

type SetCacheInput struct {
	ExpiredAt time.Time
	Payload   entity.CompanyCacheEntity
}

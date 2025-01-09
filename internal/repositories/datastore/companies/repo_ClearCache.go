package companies

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
)

func (r *repository) ClearCache(ctx context.Context) (err error) {
	prefix := "companies*"
	var cursor uint64
	var keys []string

	for {
		resp := r.redis.Do(ctx, r.redis.B().Scan().Cursor(cursor).Match(prefix).Count(100).Build())
		result, err := resp.AsScanEntry()
		if err != nil {
			return tracer.ErrInternalServer(err)
		}

		cursor = result.Cursor

		keys = append(keys, result.Elements...)

		if cursor == 0 {
			break
		}
	}

	for _, key := range keys {
		err = r.redis.Do(ctx, r.redis.B().Del().Key(key).Build()).Error()
		if err != nil {
			return tracer.ErrInternalServer(err)
		}
	}

	return
}

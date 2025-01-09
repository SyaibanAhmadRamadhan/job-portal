package infra

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisotel"
	"github.com/rs/zerolog/log"
	"strings"
)

func NewRedisWithOtel(redisConf *conf.RedisConfig, clientName string) (rueidis.Client, primitive.CloseFunc) {
	client, err := rueidisotel.NewClient(rueidis.ClientOption{
		Password:          redisConf.Password,
		ClientName:        clientName,
		InitAddress:       []string{redisConf.Host},
		CacheSizeEachConn: rueidis.DefaultCacheBytes,
	}, rueidisotel.WithDBStatement(func(cmdTokens []string) string {
		return strings.Join(cmdTokens, " | ")
	}))
	util.Panic(err)

	fn := func(ctx context.Context) error {
		log.Info().Msg("start closed redis client...")
		client.Close()
		log.Info().Msg("closed redis client successfully")
		return nil
	}

	return client, fn
}

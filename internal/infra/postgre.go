package infra

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

func NewPostgreCommand(config *conf.PostgreConfig) (*sqlx.DB, primitive.CloseFunc) {
	db, err := sqlx.Connect("postgres", config.DSN)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.DBMaxOpenConnection)
	db.SetMaxIdleConns(config.DBMaxIdleConnection)
	db.SetConnMaxLifetime(time.Second * time.Duration(config.ConnMaxLifetimeSecond))
	db.SetConnMaxIdleTime(time.Second * time.Duration(config.ConnIdleMaxLifetimeSecond))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	util.Panic(err)

	log.Info().Msg("initialization postgresql db command successfully")
	return db, func(ctx context.Context) (err error) {
		log.Info().Msg("starting close postgresql db command")

		err = db.Close()
		if err != nil {
			return err
		}

		log.Info().Msg("close postgresql db command successfully")
		return
	}
}

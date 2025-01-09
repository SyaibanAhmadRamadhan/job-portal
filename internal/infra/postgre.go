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

func NewPostgreCommand(config *conf.DatabaseConfig) (*sqlx.DB, primitive.CloseFunc) {
	db, err := sqlx.Connect("postgres", config.COMMAND.DSN)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.COMMAND.DBMaxOpenConnection)
	db.SetMaxIdleConns(config.COMMAND.DBMaxIdleConnection)
	db.SetConnMaxLifetime(time.Second * time.Duration(config.COMMAND.ConnMaxLifetimeSecond))
	db.SetConnMaxIdleTime(time.Second * time.Duration(config.COMMAND.ConnIdleMaxLifetimeSecond))

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

func NewPostgreQuery(config *conf.DatabaseConfig) (*sqlx.DB, primitive.CloseFunc) {
	db, err := sqlx.Connect("postgres", config.READER.DSN)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.READER.DBMaxOpenConnection)
	db.SetMaxIdleConns(config.READER.DBMaxIdleConnection)
	db.SetConnMaxLifetime(time.Second * time.Duration(config.READER.ConnMaxLifetimeSecond))
	db.SetConnMaxIdleTime(time.Second * time.Duration(config.READER.ConnIdleMaxLifetimeSecond))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	util.Panic(err)

	log.Info().Msg("initialization postgresql db reader successfully")
	return db, func(ctx context.Context) (err error) {
		log.Info().Msg("starting close postgresql db reader")

		err = db.Close()
		if err != nil {
			return err
		}

		log.Info().Msg("close postgresql db reader successfully")
		return
	}
}

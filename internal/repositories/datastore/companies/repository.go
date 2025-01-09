package companies

import (
	"github.com/Masterminds/squirrel"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/redis/rueidis"
)

type repository struct {
	rdbms wsqlx.Rdbms
	redis rueidis.Client
	sq    squirrel.StatementBuilderType
}

func New(rdbms wsqlx.Rdbms, redis rueidis.Client) *repository {
	return &repository{
		rdbms: rdbms,
		sq:    squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		redis: redis,
	}
}

package companies

import (
	"github.com/Masterminds/squirrel"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
)

type repository struct {
	rdbmsCommand wsqlx.Rdbms
	rdbmsReader  wsqlx.Rdbms
	sq           squirrel.StatementBuilderType
}

func New(rdbmsCommand wsqlx.Rdbms, rdbmsReader wsqlx.Rdbms) *repository {
	return &repository{
		rdbmsCommand: rdbmsCommand,
		rdbmsReader:  rdbmsReader,
		sq:           squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

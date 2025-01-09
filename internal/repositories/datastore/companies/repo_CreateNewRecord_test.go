package companies_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"math/rand"
	"regexp"
	"testing"
)

func Test_repository_CreateNewRecord(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	wsqlxDB := wsqlx.NewRdbms(sqlxDB)
	ctx := context.TODO()
	repo := companies.New(wsqlxDB, wsqlxDB)

	t.Run("should be return correct", func(t *testing.T) {
		uuid.SetRand(rand.New(rand.NewSource(1)))
		defer uuid.SetRand(nil)
		id := "52fdfc07-2182-454f-963f-5f0f9a621d72"

		expectedInput := companies.CreateNewRecordInput{
			Tx:   wsqlxDB,
			Name: "REDIKRU",
		}
		expectedOutput := companies.CreateNewRecordOutput{
			ID: id,
		}

		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO companies (id,name) VALUES ($1,$2) ON CONFLICT(name) DO UPDATE SET name=EXCLUDED.name RETURNING id`,
		)).WithArgs(id, expectedInput.Name).WillReturnRows(sqlmock.NewRows(
			[]string{"id"}).AddRow(expectedOutput.ID))

		output, err := repo.CreateNewRecord(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})
}

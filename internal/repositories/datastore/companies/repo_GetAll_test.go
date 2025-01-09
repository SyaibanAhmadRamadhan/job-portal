package companies_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func Test_repository_GetAll(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	wsqlxDB := wsqlx.NewRdbms(sqlxDB)
	ctx := context.TODO()
	repo := companies.New(wsqlxDB, nil)

	t.Run("should be return correct", func(t *testing.T) {
		expectedInput := companies.GetAllInput{
			Pagination: primitive.PaginationInput{
				Page:     1,
				PageSize: 10,
			},
		}
		expectedOutput := companies.GetAllOutput{
			Pagination: primitive.CreatePaginationOutput(expectedInput.Pagination, 3),
			Items: []entity.CompanyEntity{
				{
					ID:   uuid.NewString(),
					Name: "KREDITKRU",
				},
				{
					ID:   uuid.NewString(),
					Name: "Goto Group",
				},
				{
					ID:   uuid.NewString(),
					Name: "Shopee",
				},
			},
		}

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT count(*) FROM companies`,
		)).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT id, name FROM companies LIMIT 10 OFFSET 0`,
		)).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(expectedOutput.Items[0].ID, expectedOutput.Items[0].Name).
			AddRow(expectedOutput.Items[1].ID, expectedOutput.Items[1].Name).
			AddRow(expectedOutput.Items[2].ID, expectedOutput.Items[2].Name),
		)
		output, err := repo.GetAll(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})
}

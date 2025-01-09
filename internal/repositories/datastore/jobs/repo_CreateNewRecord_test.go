package jobs_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/jobs"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func Test_repository_CreateNewRecord(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	wsqlxDB := wsqlx.NewRdbms(sqlxDB)
	ctx := context.TODO()
	repo := jobs.New(wsqlxDB)

	t.Run("should be return correct", func(t *testing.T) {
		uuid.SetRand(rand.New(rand.NewSource(1)))
		defer uuid.SetRand(nil)
		id := "9566c74d-1003-4c4d-bbbb-0407d1e2c649"
		expectedInput := jobs.CreateNewRecordInput{
			Tx:          wsqlxDB,
			CompanyID:   uuid.NewString(),
			Title:       "Title",
			Description: "Description",
		}
		expectedOutput := jobs.CreateNewRecordOutput{
			ID:        id,
			Timestamp: time.Now().UTC(),
		}

		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO jobs (id,company_id,title,description) VALUES ($1,$2,$3,$4) RETURNING timestamp`,
		)).WithArgs(id, expectedInput.CompanyID, expectedInput.Title, expectedInput.Description).
			WillReturnRows(sqlmock.NewRows([]string{"timestamp"}).
				AddRow(expectedOutput.Timestamp))

		output, err := repo.CreateNewRecord(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})
}

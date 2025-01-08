package entity

import "time"

type JobEntity struct {
	ID          string    `db:"id" json:"id"`
	CompanyID   string    `db:"company_id" json:"company_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Timestamp   time.Time `db:"timestamp" json:"timestamp"`
}

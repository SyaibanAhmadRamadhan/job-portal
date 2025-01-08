package entity

import "time"

type IndexJobEntity struct {
	Company     IndexJobCompany `db:"company" json:"company"`
	ID          string          `db:"id" json:"id"`
	Title       string          `db:"title" json:"title"`
	Description string          `db:"description" json:"description"`
	Timestamp   time.Time       `db:"timestamp" json:"timestamp"`
}

type IndexJobCompany struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

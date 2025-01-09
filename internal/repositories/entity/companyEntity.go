package entity

import "github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"

type CompanyEntity struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type CompanyCacheEntity struct {
	Pagination primitive.PaginationOutput `json:"pagination"`
	Items      []CompanyEntity            `json:"items"`
}

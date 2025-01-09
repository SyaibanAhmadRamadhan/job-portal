package company

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
)

type service struct {
	companyRepository companies.CompanyRepository
}

type Options struct {
	CompanyRepository companies.CompanyRepository
}

func New(o Options) *service {
	return &service{
		companyRepository: o.CompanyRepository,
	}
}

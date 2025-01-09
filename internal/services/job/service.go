package job

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/index_jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
)

type service struct {
	indexJobRepository          index_jobs.JobRepository
	jobRepository               jobs.JobRepository
	companyRepository           companies.CompanyRepository
	eventBusPublisherRepository eventbus.PublisherRepository

	dbTx wsqlx.Tx
}

type Options struct {
	IndexJobRepository          index_jobs.JobRepository
	JobRepository               jobs.JobRepository
	CompanyRepository           companies.CompanyRepository
	EventBusPublisherRepository eventbus.PublisherRepository
	DBTx                        wsqlx.Tx
}

func New(o Options) *service {
	return &service{
		indexJobRepository:          o.IndexJobRepository,
		jobRepository:               o.JobRepository,
		companyRepository:           o.CompanyRepository,
		eventBusPublisherRepository: o.EventBusPublisherRepository,
		dbTx:                        o.DBTx,
	}
}

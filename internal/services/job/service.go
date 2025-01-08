package job

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/index_jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
)

type service struct {
	indexJobRepository          index_jobs.JobRepository
	jobRepository               jobs.JobRepository
	companyRepository           companies.CompanyRepository
	eventBusPublisherRepository eventbus.PublisherRepository
}

type Options struct {
	IndexJobRepository          index_jobs.JobRepository
	JobRepository               jobs.JobRepository
	CompanyRepository           companies.CompanyRepository
	EventBusPublisherRepository eventbus.PublisherRepository
}

func New(o Options) *service {
	return &service{
		indexJobRepository:          o.IndexJobRepository,
		jobRepository:               o.JobRepository,
		companyRepository:           o.CompanyRepository,
		eventBusPublisherRepository: o.EventBusPublisherRepository,
	}
}

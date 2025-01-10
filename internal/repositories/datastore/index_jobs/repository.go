package index_jobs

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type repository struct {
	client    *elasticsearch.TypedClient
	indexName string
}

func New(client *elasticsearch.TypedClient) *repository {
	return &repository{
		client:    client,
		indexName: "index_job",
	}
}

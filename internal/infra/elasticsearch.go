package infra

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/guregu/null/v5"
	"go.opentelemetry.io/otel"
	"net/http"
)

func NewES(config *conf.ElasticsearchConfig) (*elasticsearch.TypedClient, primitive.CloseFunc) {
	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses:       []string{config.Host},
		Instrumentation: elastictransport.NewOtelInstrumentation(otel.GetTracerProvider(), true, "8.10"),
		Password:        "changeme",
		Username:        "elastic",
	})
	util.Panic(err)

	_, err = typedClient.Indices.Create("index_jobs_v1").
		Request(&create.Request{
			Aliases: map[string]types.Alias{
				"index_job": {},
			},
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"company": types.ObjectProperty{
						Meta: make(map[string]string),
						Properties: map[string]types.Property{
							"id": types.KeywordProperty{
								Fields:     make(map[string]types.Property),
								Meta:       make(map[string]string),
								Properties: make(map[string]types.Property),
								DocValues:  null.BoolFrom(false).Ptr(),
							},
							"name": types.KeywordProperty{
								Fields:     make(map[string]types.Property),
								Meta:       make(map[string]string),
								Properties: make(map[string]types.Property),
								DocValues:  null.BoolFrom(false).Ptr(),
								Index:      null.BoolFrom(false).Ptr(),
							},
						},
						Dynamic: &dynamicmapping.Strict,
						Fields:  make(map[string]types.Property),
					},
					"id": types.KeywordProperty{
						Fields:     make(map[string]types.Property),
						Meta:       make(map[string]string),
						Properties: make(map[string]types.Property),
						DocValues:  null.BoolFrom(false).Ptr(),
					},
					"title":       types.NewTextProperty(),
					"description": types.NewTextProperty(),
					"timestamp":   types.NewDateProperty(),
				},
			},
		}).
		Do(nil)

	//util.Panic(err)
	if err != nil {
		var esError *types.ElasticsearchError
		if errors.As(err, &esError) {
			if esError.Status == http.StatusBadRequest {
				return typedClient, func(ctx context.Context) error {
					return nil
				}
			}
			panic(err)
		}
	}

	return typedClient, func(ctx context.Context) error {
		return nil
	}
}

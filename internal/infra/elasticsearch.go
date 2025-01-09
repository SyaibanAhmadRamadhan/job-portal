package infra

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func NewES(config *conf.ElasticsearchConfig) {
	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{config.Host},
	})
	util.Panic(err)

	res, err := typedClient.Indices.Create("index_jobs_v1").
		Request(&create.Request{
			Aliases: map[string]types.Alias{
				"index_job": {},
			},
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"company": types.ObjectProperty{
						Fields: map[string]types.Property{
							"id": types.KeywordProperty{
								Boost:                    nil,
								CopyTo:                   nil,
								DocValues:                nil,
								Dynamic:                  nil,
								EagerGlobalOrdinals:      nil,
								Fields:                   nil,
								IgnoreAbove:              nil,
								Index:                    nil,
								IndexOptions:             nil,
								Meta:                     nil,
								Normalizer:               nil,
								Norms:                    nil,
								NullValue:                nil,
								OnScriptError:            nil,
								Properties:               nil,
								Script:                   nil,
								Similarity:               nil,
								SplitQueriesOnWhitespace: nil,
								Store:                    nil,
								TimeSeriesDimension:      nil,
								Type:                     "",
							},
						},
						Properties: make(map[string]types.Property),
					},
				},
			},
			Settings: nil,
		}).
		Do(nil)
}

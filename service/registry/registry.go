package registry

import (
	"github.com/hardeepnarang10/query/pkg/authorization"
	"github.com/hardeepnarang10/query/pkg/elasticsearch"
)

type ServiceRegistry struct {
	ServiceConfig       ServiceConfig
	AuthorizationClient authorization.Authorization
	ElasticsearchClient elasticsearch.ElasticSearch
}

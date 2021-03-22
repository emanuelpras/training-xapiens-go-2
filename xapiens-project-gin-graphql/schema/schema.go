package schema

import (
	"xapiens-project-gin-graphql/api"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    api.QueryType,
		Mutation: api.MutationType,
	},
)

package root

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-subscription-demo/queries"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping":  queries.PingQuery,
		"hello": queries.HelloQuery,
	},
})

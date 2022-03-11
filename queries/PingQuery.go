package queries

import "github.com/graphql-go/graphql"

var PingQuery = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "ok", nil
	},
}

package mutations

import "github.com/graphql-go/graphql"

var CreateUserMutation = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "Oui", nil
	},
}

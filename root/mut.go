package root

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-subscription-demo/mutations"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": mutations.CreateUserMutation,
	},
})

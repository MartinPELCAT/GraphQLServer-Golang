package server

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-subscription-demo/root"
	"log"
)

func GetSchema() graphql.Schema {
	schemaConfig := graphql.SchemaConfig{
		Query:        root.RootQuery,
		Subscription: root.RootSubscription,
		Mutation:     root.RootMutation,
	}

	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal(err)
	}

	return s
}

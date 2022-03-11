package main

import (
	"context"
	"github.com/graphql-go/graphql-go-subscription-demo/server"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	s := server.GetSchema()
	h := handler.New(&handler.Config{
		Schema:     &s,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	http.HandleFunc("/gql", func(writer http.ResponseWriter, request *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "test", "test")
		h.ContextHandler(ctx, writer, request)
	})
	http.HandleFunc("/subscriptions", server.SubscriptionsHandler)

	log.Println("GraphQL Server running on [POST]: localhost:8081/graphql")
	log.Println("GraphQL Playground running on [GET]: localhost:8081/graphql")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

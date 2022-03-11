package root

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-subscription-demo/subscriptions"
)

var RootSubscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"feed": subscriptions.FeedSubscription,
	},
})

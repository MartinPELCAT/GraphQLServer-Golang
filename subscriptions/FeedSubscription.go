package subscriptions

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-subscription-demo/objects"
	"log"
	"time"
)

var FeedSubscription = &graphql.Field{
	Type: objects.FeedType,
	Args: map[string]*graphql.ArgumentConfig{
		"to": {
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return p.Source, nil
	},
	Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
		c := make(chan interface{})
		to := p.Args["to"].(int)

		go func() {
			var i int

			for {
				i++

				feed := objects.Feed{ID: fmt.Sprintf("%d", i)}

				select {
				case <-p.Context.Done():
					log.Println("[RootSubscription] [Subscribe] subscription canceled")
					close(c)
					return
				default:
					c <- feed
				}

				time.Sleep(250 * time.Millisecond)

				if i == to {
					close(c)
					return
				}
			}
		}()

		return c, nil
	},
}

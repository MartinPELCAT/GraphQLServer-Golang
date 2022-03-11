package queries

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var HelloQuery = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		test := p.Context.Value("test")

		var oui = fmt.Sprintf("%v", test)

		return "World !" + oui, nil
	},
}

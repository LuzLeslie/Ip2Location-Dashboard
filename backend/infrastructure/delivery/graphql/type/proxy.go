package ipInfo

import (
	"github.com/graphql-go/graphql"
)

var ProxyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Proxy",
		Fields: graphql.Fields{
			"is_proxy": &graphql.Field{
				Type: graphql.Int,
			},
			"proxy_type": &graphql.Field{
				Type: graphql.Int,
			},
			// "percentage": &graphql.Field{
			// 	Type: graphql.Int,
			// },
		},
	},
)

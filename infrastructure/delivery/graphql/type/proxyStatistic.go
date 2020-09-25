package ipInfo

import (
	"github.com/graphql-go/graphql"
)

var ProxyStatistic = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProxyStatistic",
		Fields: graphql.Fields{
			"proxy": &graphql.Field{
				Type: graphql.Int,
			},
			"clean": &graphql.Field{
				Type: graphql.Int,
			},
			// "percentage": &graphql.Field{
			// 	Type: graphql.Int,
			// },
		},
	},
)

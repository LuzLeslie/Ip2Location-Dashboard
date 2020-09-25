package ipInfo

import (
	"github.com/graphql-go/graphql"
)

var TopCountrysType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TopCountrys",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
			// "percentage": &graphql.Field{
			// 	Type: graphql.Int,
			// },
		},
	},
)

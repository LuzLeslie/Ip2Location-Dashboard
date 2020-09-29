package graphql

import (
	ipInfo "github.com/LuzLeslie/Ip2Location-Dashboard/infrastructure/delivery/graphql/type"
	"github.com/graphql-go/graphql"
)

func (s Schema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"analyze": &graphql.Field{
				Type:        ipInfo.IpType,
				Description: "ip to analyze",
				Args: graphql.FieldConfigArgument{
					"ip": &graphql.ArgumentConfig{
						Description: "ip",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: s.AnalyzeIp,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

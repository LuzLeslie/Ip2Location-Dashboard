package graphql

import (
	ipInfo "github.com/LuzLeslie/Ip2Location-Dashboard/infrastructure/delivery/graphql/type"
	"github.com/graphql-go/graphql"
	// log "github.com/sirupsen/logrus"
)

func (s Schema) Subscription() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			"ipData": &graphql.Field{
				Type:        ipInfo.IpType,
				Description: "ip to analyze",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return s.AnalyzeUseCase.GetLastAnalyzed()
				},
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

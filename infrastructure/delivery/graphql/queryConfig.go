package graphql

import (
	types "github.com/LuzLeslie/Ip2Location-Dashboard/infrastructure/delivery/graphql/type"
	"github.com/graphql-go/graphql"
)

func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"ipData": &graphql.Field{
				Type:        graphql.NewList(types.IpType),
				Description: "Get all log ip",
				Resolve:     s.GetIpLog,
			},
			"topCountrys": &graphql.Field{
				Type:        graphql.NewList(types.TopCountrysType),
				Description: "Get statistics",
				Resolve:     s.GetTopCountrys,
			},
			"proxyStatistic": &graphql.Field{
				Type:        types.ProxyStatistic,
				Description: "Get percentage",
				Resolve:     s.GetProxyStatistic,
			},
			"totalAnalyzed": &graphql.Field{
				Type:        graphql.Int,
				Description: "Number of ip analyzed",
				Resolve:     s.GetTotalAnalyzed,
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

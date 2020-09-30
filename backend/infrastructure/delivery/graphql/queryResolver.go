package graphql

import (
	"github.com/graphql-go/graphql"
	// log "github.com/sirupsen/logrus"
)

func (s Schema) GetIpLog(params graphql.ResolveParams) (interface{}, error) {
	return s.AnalyzeUseCase.GetLog()
}

func (s Schema) GetTopCountrys(params graphql.ResolveParams) (interface{}, error) {
	return s.AnalyzeUseCase.GetTopCountrys()
}

func (s Schema) GetProxyStatistic(params graphql.ResolveParams) (interface{}, error) {
	return s.AnalyzeUseCase.GetProxyStatistic()
}

func (s Schema) GetTotalAnalyzed(params graphql.ResolveParams) (interface{}, error) {
	return s.AnalyzeUseCase.GetTotalAnalyzed()
}

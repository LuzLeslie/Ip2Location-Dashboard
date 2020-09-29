package graphql

import (
	analyzeuc "github.com/LuzLeslie/Ip2Location-Dashboard/package/analyze/usecase"
	"github.com/graphql-go/graphql"
	// log "github.com/sirupsen/logrus"
)

type Schema struct {
	AnalyzeUseCase analyzeuc.AnalyzeUseCase
	RootSchema     *graphql.Schema
}

func New(analyzeUC *analyzeuc.AnalyzeUseCase) *Schema {
	return &Schema{
		AnalyzeUseCase: *analyzeUC,
	}
}

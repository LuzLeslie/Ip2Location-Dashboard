package repository

import (
	models "github.com/LuzLeslie/Ip2Location-Dashboard/domain/models"
	simpleDB "github.com/LuzLeslie/Ip2Location-Dashboard/infrastructure/datastore/simpleDB"
	// log "github.com/sirupsen/logrus"
)

type AnalyzeRepo struct {
	db simpleDB.SimpleDB
}

func New(db *simpleDB.SimpleDB) *AnalyzeRepo {
	return &AnalyzeRepo{
		db: *db,
	}
}

func (conn *AnalyzeRepo) SaveLog(infoIp *models.Ip) error {
	conn.db.InsertLog(infoIp)
	return nil
}

func (conn *AnalyzeRepo) GetLog() ([]models.Ip, error) {
	return conn.db.GetLog()
}

func (conn *AnalyzeRepo) GetLastAnalyzed() (models.Ip, error) {
	return conn.db.GetLastAnalyzed()
}

func (conn *AnalyzeRepo) GetTopCountrys() ([]models.TopCountrys, error) {
	return conn.db.GetTopCountrys()
}

func (conn *AnalyzeRepo) AddNewTopCountry(newCountry models.TopCountrys) {
	conn.db.InsertNewTopCountry(newCountry)
}

func (conn *AnalyzeRepo) IncreaseTopCountry(idCountry int) {
	conn.db.IncreaseTopCountry(idCountry)
}

func (conn *AnalyzeRepo) GetTotalAnalyzed() (int, error) {
	return conn.db.GetTotalAnalyzed(), nil
}

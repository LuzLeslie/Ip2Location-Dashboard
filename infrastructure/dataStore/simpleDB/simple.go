package simpleDB

import (
	models "github.com/LuzLeslie/ip-debugger/domain/models"
	// log "github.com/sirupsen/logrus"
)

type SimpleDB struct {
	logIp       []models.Ip
	topCountrys []models.TopCountrys
}

func New() *SimpleDB {
	return &SimpleDB{
		logIp:       []models.Ip{},
		topCountrys: []models.TopCountrys{},
	}
}

func (s *SimpleDB) InsertLog(infoIp *models.Ip) {
	s.logIp = append(s.logIp, *infoIp)
}

func (s *SimpleDB) GetLog() ([]models.Ip, error) {
	return s.logIp, nil
}

func (s *SimpleDB) GetLastAnalyzed() (models.Ip, error) {
	return s.logIp[len(s.logIp)-1], nil
}

func (s *SimpleDB) GetTopCountrys() ([]models.TopCountrys, error) {
	return s.topCountrys, nil
}

func (s *SimpleDB) InsertNewTopCountry(newCountry models.TopCountrys) error {
	s.topCountrys = append(s.topCountrys, newCountry)
	return nil
}

func (s *SimpleDB) IncreaseTopCountry(pos int) {
	s.topCountrys[pos].Quantity = s.topCountrys[pos].Quantity + 1
}

func (s *SimpleDB) GetTotalAnalyzed() int {
	return len(s.logIp)
}

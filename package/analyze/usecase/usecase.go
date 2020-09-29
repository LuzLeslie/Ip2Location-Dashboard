package usecase

import (
	models "github.com/LuzLeslie/Ip2Location-Dashboard/domain/models"
	analyzeRepo "github.com/LuzLeslie/Ip2Location-Dashboard/package/analyze/repository"
	locationUseCase "github.com/LuzLeslie/Ip2Location-Dashboard/package/location/usecase"
	proxyUseCase "github.com/LuzLeslie/Ip2Location-Dashboard/package/proxy/usecase"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type AnalyzeUseCase struct {
	analyzeRepo   *analyzeRepo.AnalyzeRepo
	LocationUcase locationUseCase.LocationUseCase
	ProxyUcase    proxyUseCase.ProxyUseCase
}

func New(repo *analyzeRepo.AnalyzeRepo, locationUsecase *locationUseCase.LocationUseCase, proxyUsecase *proxyUseCase.ProxyUseCase) *AnalyzeUseCase {
	return &AnalyzeUseCase{
		analyzeRepo:   repo,
		LocationUcase: *locationUsecase,
		ProxyUcase:    *proxyUsecase,
	}
}

func (a *AnalyzeUseCase) GetInfoIp(ip string) (*models.Ip, error) {
	info, err := a.LocationUcase.GetInfo(ip)

	log.Debug(info)

	if err != nil {
		return nil, err
	}
	info.Ip = ip
	result, err := a.ProxyUcase.GetInfo(ip)
	if err != nil {
		log.Error("Proxy error result: ", err)
	} else {
		info.Proxy = *result
	}

	a.analyzeRepo.SaveLog(info)
	a.updateTopCountrys(info.Country_name)

	return info, nil
}

func (a *AnalyzeUseCase) updateTopCountrys(countryName string) {
	currentTopCountrys, _ := a.analyzeRepo.GetTopCountrys()
	possibleId := getIdCountry(currentTopCountrys, countryName)

	if possibleId != -1 {
		a.analyzeRepo.IncreaseTopCountry(possibleId)
	} else {
		newCountry := models.TopCountrys{
			Name:     countryName,
			Quantity: 1,
		}
		a.analyzeRepo.AddNewTopCountry(newCountry)
	}
}

func getIdCountry(currentTopCountrys []models.TopCountrys, countryName string) int {
	for i, country := range currentTopCountrys {
		if country.Name == countryName {
			return i
		}
	}
	return -1
}

func (a *AnalyzeUseCase) GetTopCountrys() ([]models.TopCountrys, error) {
	return a.analyzeRepo.GetTopCountrys()
}

func (a *AnalyzeUseCase) GetLog() ([]models.Ip, error) {
	return a.analyzeRepo.GetLog()
}

func (a *AnalyzeUseCase) GetLastAnalyzed() (models.Ip, error) {
	return a.analyzeRepo.GetLastAnalyzed()
}

func (a *AnalyzeUseCase) GetProxyStatistic() (models.ProxyStatistic, error) {
	allLog, err := a.analyzeRepo.GetLog()
	if err != nil {
		log.Fatal(err)
	}
	proxyCount := 0

	for _, proxy := range allLog {
		isProxy := proxy.Proxy.Is_proxy
		number, err := strconv.Atoi(isProxy)
		if err != nil {
			log.Error(err)
		}
		if number > 0 {
			proxyCount += 1
		}
	}

	percentageProxy := proxyCount * 100 / len(allLog)

	proxyStat := models.ProxyStatistic{
		Proxy: percentageProxy,
		Clean: 100 - percentageProxy,
	}
	return proxyStat, nil
}

func (a *AnalyzeUseCase) GetTotalAnalyzed() (int, error) {
	return a.analyzeRepo.GetTotalAnalyzed()
}

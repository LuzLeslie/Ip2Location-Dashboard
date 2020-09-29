package usecase

import (
	models "github.com/LuzLeslie/Ip2Location-Dashboard/domain/models"
	locationRepository "github.com/LuzLeslie/Ip2Location-Dashboard/package/location/repository"
	// log "github.com/sirupsen/logrus"
)

type LocationUseCase struct {
	LocationRepo locationRepository.LocationRepo
}

func New(repo *locationRepository.LocationRepo) *LocationUseCase {
	return &LocationUseCase{
		LocationRepo: *repo,
	}
}

func (l *LocationUseCase) GetInfo(ip string) (*models.Ip, error) {
	return l.LocationRepo.GetInfo(ip)
}

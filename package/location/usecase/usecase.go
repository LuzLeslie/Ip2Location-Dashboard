package usecase

import (
	models "github.com/LuzLeslie/ip-debugger/domain/models"
	locationRepository "github.com/LuzLeslie/ip-debugger/package/location/repository"
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

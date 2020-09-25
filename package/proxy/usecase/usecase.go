package usecase

import (
	models "github.com/LuzLeslie/ip-debugger/domain/models"
	proxyRepository "github.com/LuzLeslie/ip-debugger/package/proxy/repository"
	// log "github.com/sirupsen/logrus"
)

type ProxyUseCase struct {
	ProxyRepo proxyRepository.ProxyRepo
}

func New(repoProxy *proxyRepository.ProxyRepo) *ProxyUseCase {
	return &ProxyUseCase{
		ProxyRepo: *repoProxy,
	}
}

func (p *ProxyUseCase) GetInfo(ip string) (*models.Proxy, error) {
	return p.ProxyRepo.GetInfo(ip)
}

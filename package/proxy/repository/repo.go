package repository

import (
	"errors"
	// "fmt"
	models "github.com/LuzLeslie/ip-debugger/domain/models"
	"github.com/ip2location/ip2proxy-go"
	log "github.com/sirupsen/logrus"
)

type ProxyRepo struct {
	pathBin string
}

func New(pathBin string) *ProxyRepo {
	return &ProxyRepo{
		pathBin: pathBin,
	}
}

func (p *ProxyRepo) GetInfo(ip string) (*models.Proxy, error) {

	dbIp2Proxy, err := ip2proxy.OpenDB(p.pathBin)

	if err != nil {
		log.Fatal("Database proxy open failed: ", err)
		return nil, errors.New("Database open failed")
	}

	result, err := dbIp2Proxy.GetAll(ip)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	proxyInfo := models.Proxy{
		Is_proxy:   result["isProxy"],
		Proxy_type: result["ProxyType"],
	}

	defer dbIp2Proxy.Close()

	return &proxyInfo, nil
}

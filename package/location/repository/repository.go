package repository

import (
	"errors"
	// "fmt"
	models "github.com/LuzLeslie/ip-debugger/domain/models"
	"github.com/ip2location/ip2location-go"
	log "github.com/sirupsen/logrus"
)

type LocationRepo struct {
	LocationBinDB string
}

func New(locationBinDBx string) *LocationRepo {
	return &LocationRepo{
		LocationBinDB: locationBinDBx,
	}
}

func (l *LocationRepo) GetInfo(ip string) (*models.Ip, error) {
	dbIp2Location, err := ip2location.OpenDB(l.LocationBinDB)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Database open failed")
	}

	resultsLocation, err := dbIp2Location.Get_all(ip)

	if err != nil {
		return nil, err
	}

	// log.Debug(resultsLocation)

	ipInfo := models.Ip{
		Country_code:         resultsLocation.Country_short,
		Country_name:         resultsLocation.Country_long,
		Region:               resultsLocation.Region,
		City:                 resultsLocation.City,
		Isp:                  resultsLocation.Isp,
		Latitude:             resultsLocation.Latitude,
		Longitude:            resultsLocation.Longitude,
		Domain:               resultsLocation.Domain,
		Zip_code:             resultsLocation.Zipcode,
		Time_zone:            resultsLocation.Timezone,
		Net_speed:            resultsLocation.Netspeed,
		Idd_code:             resultsLocation.Iddcode,
		Area_code:            resultsLocation.Areacode,
		Weather_station_code: resultsLocation.Weatherstationcode,
		Weather_station_name: resultsLocation.Weatherstationname,
		Mcc:                  resultsLocation.Mcc,
		Mnc:                  resultsLocation.Mnc,
		Mobile_brand:         resultsLocation.Mobilebrand,
		Elevation:            resultsLocation.Elevation,
		// Usage_type:           resultsLocation.Usagetype,
	}

	defer dbIp2Location.Close()

	return &ipInfo, nil
}

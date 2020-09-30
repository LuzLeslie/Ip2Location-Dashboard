package models

type Ip struct {
	Ip           string `json:"ip"`
	Country_code string `json:"country_code"`
	Country_name string `json:"country_name"`
	// Count_request      int    `json:"count_request"`
	Region               string  `json:"region"`
	City                 string  `json:"city"`
	Isp                  string  `json:"isp"`
	Latitude             float32 `json:"latitude"`
	Longitude            float32 `json:"longitude"`
	Domain               string  `json:"domain"`
	Zip_code             string  `json:"zip_code"`
	Time_zone            string  `json:"time_zone"`
	Net_speed            string  `json:"nets_peed"`
	Idd_code             string  `json:"idd_code"`
	Area_code            string  `json:"area_code"`
	Weather_station_code string  `json:"weather_station_code"`
	Weather_station_name string  `json:"weather_station_name"`
	Mcc                  string  `json:"mcc"`
	Mnc                  string  `json:"mnc"`
	Mobile_brand         string  `json:"mobile_brand"`
	Elevation            float32 `json:"elevation"`
	Proxy                Proxy   `json:"proxy"`
}

type Chat struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Message string `json:"message"`
}

type Country struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Statistics struct {
	Countrys []Country `json:"countrys"`
}

package ipInfo

import (
	"github.com/graphql-go/graphql"
)

var IpType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IpDetails",
		Fields: graphql.Fields{
			"ip": &graphql.Field{
				Type: graphql.String,
			},
			"country_code": &graphql.Field{
				Type: graphql.String,
			},
			"country_name": &graphql.Field{
				Type: graphql.String,
			},
			"region": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"isp": &graphql.Field{
				Type: graphql.String,
			},
			"latitude": &graphql.Field{
				Type: graphql.String,
			},
			"longitude": &graphql.Field{
				Type: graphql.String,
			},
			"domain": &graphql.Field{
				Type: graphql.String,
			},
			"zip_code": &graphql.Field{
				Type: graphql.String,
			},
			"time_zone": &graphql.Field{
				Type: graphql.String,
			},
			"net_speed": &graphql.Field{
				Type: graphql.String,
			},
			"idd_code": &graphql.Field{
				Type: graphql.String,
			},
			"area_code": &graphql.Field{
				Type: graphql.String,
			},
			"weather_station_code": &graphql.Field{
				Type: graphql.String,
			},
			"weather_station_name": &graphql.Field{
				Type: graphql.String,
			},
			"mcc": &graphql.Field{
				Type: graphql.String,
			},
			"mnc": &graphql.Field{
				Type: graphql.String,
			},
			"mobile_brand": &graphql.Field{
				Type: graphql.String,
			},
			"elevation": &graphql.Field{
				Type: graphql.String,
			},
			"usage_type": &graphql.Field{
				Type: graphql.String,
			},
			"proxy": &graphql.Field{
				Type: ProxyType,
			},

			// "count_request": &graphql.Field{
			// 	Type: graphql.Int,
			// },
		},
	},
)

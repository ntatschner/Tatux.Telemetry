package db

import (
	influxdb2 "influxdata/influxdb-client-go/v2"
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
)



func NewInfluxDBClient(url string, token string) influxdb2.Client {
	return influxdb2.NewClient(url, token)
}


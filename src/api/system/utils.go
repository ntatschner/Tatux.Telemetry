package system

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-io-go/ip2locationio"
)

var (
	ip2location_api_key = GetEnv("IP2LOCATION_API_KEY", "", false)
)

func GetEnv(key, defaultValue string, throwOnDefault bool) string {
	value, exists := os.LookupEnv(key)
	if !exists && !throwOnDefault {
		return defaultValue
	} else if !exists && throwOnDefault {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func GetClientIP(c *gin.Context) string {
	clientIP := c.ClientIP()
	if clientIP == "" {
		clientIP = c.Request.Header.Get("X-Forwarded-For")
	}
	return clientIP
}

func GetGeoLocation(ip string) (float64, float64) {
	config, err := ip2locationio.OpenConfiguration(ip2location_api_key)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return 0, 0
	}

	ipl, err := ip2locationio.OpenIPGeolocation(config)

	if err != nil {
		log.Fatalf("Failed to get config.")
		return 0, 0
	}

	lang := ""
	results, err := ipl.LookUp(ip, lang)

	if err != nil {
		log.Fatalf("Failed to lookup ip: %s`nError: %t", ip, err)
		return 0, 0
	}

	return results.Latitude, results.Longitude
}

package system

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
		log.Fatalf("Failed to load geolocation database file: %v", err)
		return 0, 0
	}

	ipl, err := ip2locationio.OpenIPGeolocation(config)

	if err != nil {
		fmt.Print(err)
		return 0, 0
	}

	lang := ""
	results, err := ipl.LookUp(ip, lang)

	if err != nil {
		fmt.Print(err)
		return 0, 0
	}

	return results.Latitude, results.Longitude
}

func GetGeoLocationDatabase() {
	url := "https://www.ip2location.com/download/?token=vH6gLcMcVBMaibeswIowRFCcbWXsWSGCHeXFxauF5RIMdruzYTVTCzgn6BTHOx21&file=DB5LITEBIN"
	fmt.Println("Downloading Geo Location File.")
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to Download Geo Location File: %v", err)
	}
	defer response.Body.Close()

	file, err := os.Create("DB5LITE.BIN")
	if err != nil {
		log.Fatalf("Failed to Create Geo Location File: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Failed to Copy Geo Location File: %v", err)
	}

	fmt.Println("Geo Location File downloaded.")
}

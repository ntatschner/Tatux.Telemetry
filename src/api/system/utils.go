package system

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go/v9"
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

func GetGeoLocation(ip string) (float32, float32) {
	db, err := ip2location.OpenDB("geodata.bin")
	if err != nil {
		log.Fatalf("Failed to load geolocation database file: %v", err)
		return 0, 0
	}
	defer db.Close()

	results, err := db.Get_all(ip)
	if err != nil {
		log.Fatalf("Failed to get geolocation for IP address %s: %v", ip, err)
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

	file, err := os.Create("geodata.bin")
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

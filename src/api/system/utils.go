package system

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
		log.Fatal(err)
	}
	defer db.Close()

	results, err := db.Get_all(ip)

	if err != nil {
		fmt.Print(err)
		return 0, 0
	}
	return results.Latitude, results.Longitude
}

func GetGeoLocationDatabase() {
	url := "https://www.ip2location.com/download/?token=vH6gLcMcVBMaibeswIowRFCcbWXsWSGCHeXFxauF5RIMdruzYTVTCzgn6BTHOx21&file=DB5LITEBIN"
	println("Downloading Geo Location File.")
	response, err := http.Get(url)
	if err != nil {
		println("Failed to Download Geo Location File.")
	}
	defer response.Body.Close()

	file, err := os.Create("geodata.bin")
	if err != nil {
		println("Failed to Create Geo Location File.")
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		println("Failed to Copy Geo Location File.")
	}

	println("Geo Location File downloaded.")
	time.Sleep(24 * time.Hour)

}

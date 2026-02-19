package geolocation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Sh4Ryuu/go-scan/pkg/models"
)

const (
	ipAPIURL = "http://ip-api.com/json/"
	timeout  = 10 * time.Second
)

// LookupIP looks up geolocation information for an IP address
func LookupIP(ip string) *models.GeoLocation {
	client := &http.Client{Timeout: timeout}
	url := fmt.Sprintf("%s%s?fields=status,message,country,countryCode,region,city,lat,lon,isp", ipAPIURL, ip)

	resp, err := client.Get(url)
	if err != nil {
		return &models.GeoLocation{
			IP:    ip,
			Error: err.Error(),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &models.GeoLocation{
			IP:    ip,
			Error: err.Error(),
		}
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return &models.GeoLocation{
			IP:    ip,
			Error: err.Error(),
		}
	}

	if status, ok := result["status"].(string); !ok || status != "success" {
		return &models.GeoLocation{
			IP:    ip,
			Error: "lookup failed",
		}
	}

	lat, _ := result["lat"].(float64)
	lon, _ := result["lon"].(float64)

	country, _ := result["country"].(string)
	countryCode, _ := result["countryCode"].(string)
	region, _ := result["region"].(string)
	city, _ := result["city"].(string)
	isp, _ := result["isp"].(string)

	return &models.GeoLocation{
		IP:          ip,
		Country:     country,
		CountryCode: countryCode,
		Region:      region,
		City:        city,
		Latitude:    lat,
		Longitude:   lon,
		ISP:         isp,
	}
}

// BatchLookup performs geolocation lookup for multiple IPs
func BatchLookup(ips []string) []*models.GeoLocation {
	results := make([]*models.GeoLocation, len(ips))
	for i, ip := range ips {
		results[i] = LookupIP(ip)
	}
	return results
}

package geolocation

import (
	"fmt"
	"testing"
)

const ip = "8.8.8.8"

func TestGeoLocation(t *testing.T) {
	country, countryCode, region, regionCode, city, tz, lat, lon, err := GeoLocation(ip)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("TEST RESULTS: country %s (%s), region %s (%s) city %s tz %s lat/lon %f/%f\n",
		country, countryCode, region, regionCode, city, tz, lat, lon)
}

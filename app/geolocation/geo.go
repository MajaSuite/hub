package geolocation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gohub/app"
)

var server = "http://ip-api.com/json/"

// return:
//	country, countryCode, region, regionCode, city, tz, lat, lon
func GeoLocation (ipaddr string) (string, string, string, string, string, string, float32, float32, error) {
	type JsonStruct struct {
		Status      string  `json:"status"`// "success"
		Country     string  `json:"country"`// "Russia"
		CountryCode string  `json:"countryCode"`// "RU"
		Region      string  `json:"region"`	// "MOW",
		RegionName  string  `json:"regionName"`// "Moscow"
		City        string  `json:"city"`	// "Moscow"
		Zip         string  `json:"zip"`	// "105425"
		Lat         float32 `json:"lat"`	// 55.7482
		Lon         float32 `json:"lon"`	// 37.6177
		Timezone    string  `json:"timezone"`	// "Europe/Moscow"
		Isp         string  `json:"isp"`	// ""
		Org         string  `json:"org"`	// ""
		As          string  `json:"as"`		// "AS42610 PJSC Rostelecom"
		Query       string  `json:"query"`	// "109.173.98.57"
	}
	var j JsonStruct

	resp, err := http.Get(server + ipaddr)
	if err != nil {
		return "", "", "", "", "", "", 0, 0, app.ErrNotConnected
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", "", "", 0, 0, err
	}

	err = json.Unmarshal(buf, &j)
	if err != nil {
		return "", "", "", "", "", "", 0, 0, app.ErrNotParsed
	}

	return j.Country, j.CountryCode, j.RegionName, j.Region, j.City, j.Timezone, j.Lat, j.Lon, nil
}

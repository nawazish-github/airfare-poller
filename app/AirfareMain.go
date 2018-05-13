package main

import (
	"fmt"

	"github.com/nawazish-github/airfare-poller/http"
	"github.com/nawazish-github/airfare-poller/unmarshal"
)

func main() {
	var airfareHttpClient http.AirfareHttpClient
	url := "http://developer.goibibo.com/api/search/?app_id=******&app_key=******************&format=json&source=BLR&destination=IXR&dateofdeparture=20180614&dateofarrival=20180624&seatingclass=E&adults=1&children=0&infants=0&counter=100"
	data, urlError := airfareHttpClient.Get(url)

	if urlError != nil {
		fmt.Println(urlError)
		return
	}
	var jsonUnmarshaller unmarshal.JsonUnmarshaller
	airfareResp, airfareRespError := jsonUnmarshaller.Unmarshal(data)

	if airfareRespError != nil {
		fmt.Println(airfareRespError)
		return
	}

	fmt.Println(airfareResp)
}

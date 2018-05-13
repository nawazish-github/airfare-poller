package main

import (
	"fmt"

	"github.com/nawazish-github/airfare-poller/config"
	"github.com/nawazish-github/airfare-poller/http"
	"github.com/nawazish-github/airfare-poller/unmarshal"
)

func main() {

	configFile := "./config/config.json"
	var configUnmarshaller config.ConfigUnmarshaller
	config, configError := configUnmarshaller.Unmarshal(configFile)
	if configError != nil {
		fmt.Println(configError)
	}
	var airfareHttpClient http.AirfareHttpClient
	url := "http://developer.goibibo.com/api/search/?app_id=" + config.App_id + "&app_key=" + config.App_key +
		"&format=" + config.Format + "&source=" + config.Source + "&destination=" + config.Destination +
		"&dateofdeparture=" + config.Dateofdeparture + "&dateofarrival=" + config.Dateofarrival +
		"&seatingclass=E&adults=1&children=0&infants=0&counter=100"

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

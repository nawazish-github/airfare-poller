package main

import (
	"fmt"
	"log"

	"github.com/nawazish-github/airfare-poller/database"
	"github.com/nawazish-github/airfare-poller/models"

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

	log.Output(0, "Successfully read the Config file!")
	url := getURL(config)
	var airfareHttpClient http.AirfareHttpClient
	data, urlError := airfareHttpClient.Get(url)

	if urlError != nil {
		fmt.Println(urlError)
		return
	}
	log.Output(0, "Successfully queried the remote service!")
	var jsonUnmarshaller unmarshal.JsonUnmarshaller
	airfareResp, airfareRespError := jsonUnmarshaller.Unmarshal(data)

	if airfareRespError != nil {
		fmt.Println(airfareRespError)
		return
	}
	//fmt.Println(airfareResp)
	log.Output(0, "Successfully unmarshalled the remote response!")

	/* get data from DB to compare;
	   if DB has no data, simply insert all data into it;
	   if DB has data in it, update (upsert) delta into it*/

	var databaseClient database.DatabaseClient
	_, dbConErr := databaseClient.DialDBAt("localhost", "airfare-poller-1", "airfare_poller_collection_one")
	if dbConErr != nil {
		panic(dbConErr)
	}

	log.Output(0, "Successfully dialed into database server!")
	response, dbDataErr := databaseClient.GetDataFor("BLR-IXR")
	if dbDataErr != nil {
		panic(dbDataErr)
	}

	log.Output(100, "Successfully queried Database for archived data!")
	if len(response) == 0 {
		dbUpdateError := databaseClient.UpdateDataFor("BLR-IXR", airfareResp)
		if dbUpdateError != nil {
			panic(dbUpdateError)
		}
		log.Output(0, "database updated successfully with new data...enjoy")
	}

	log.Output(0, "verifying delta in two data sets...")

	/*compare deltas between http airfareresponse and db airfareresponse
	  and*/

}

func getURL(config *models.Config) string {
	return "http://developer.goibibo.com/api/search/?app_id=" + config.App_id + "&app_key=" + config.App_key +
		"&format=" + config.Format + "&source=" + config.Source + "&destination=" + config.Destination +
		"&dateofdeparture=" + config.Dateofdeparture + "&dateofarrival=" + config.Dateofarrival +
		"&seatingclass=E&adults=1&children=0&infants=0&counter=100"
}

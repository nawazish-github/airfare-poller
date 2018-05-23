package models

//Response ...
type Response struct {
	SrcDest string `json:srcdest`
}

//AirfareResponse ...the response we get after quering end point
type AirfareResponse struct {
	SrcDest string `json:srcdest` //key of the document
	Data    Data   `json:data`
}

type Data struct {
	Returnflights []Returnflights `json:returnflights`
	Onwardflights []Returnflights `json:onwardflights`
}

type Returnflights struct {
	Origin           string `json:origin`
	Flightcode       string `json:flightcode`
	Duration         string `json:duration`
	Flightno         string `json:flightno`
	Destination      string `json:destination`
	Stops            string `json:stops`
	Seatavailable    string `json:seatavailable`
	Airline          string `json:airline`
	Fare             Fare   `json:fare` //?
	BookabilityValue float64
	Onwardflights    []Returnflights `json:onwardflights`
	Splitduration    string          `json:splitduration`
	Depdate          string          `json:depdate`
	Arrtime          string          `json:arrtime`
	Arrdate          string          `json:arrdate`
}

type Fare struct {
	Grossamount    int `json:grossamount`
	Totalbasefare  int `json:totalbasefare`
	Adultbasefare  int `json:adultbasefare`
	Totalfare      int `json:totalfare`
	Totalsurcharge int `json:totalsurcharge`
	Totaltaxes     int `json:totaltaxes`
	Adulttax       int `json:adulttax`
	Adulttotalfare int `json:adulttotalfare`
	//Totalcommission string `json:totalcommission`
}

type Config struct {
	App_id          string `json:app_id`
	App_key         string `json:app_key`
	Format          string `json:format`
	Source          string `json:source`
	Destination     string `json:destination`
	Dateofdeparture string `json:dateofdeparture`
	Dateofarrival   string `json:dateofarrival`
}

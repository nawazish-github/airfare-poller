package unmarshal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nawazish-github/airfare-poller/models"
)

type JsonUnmarshaller struct{}

func (jsonUnmarshaller JsonUnmarshaller) Unmarshal(data []byte) (*models.AirfareResponse, error) {
	url := "http://developer.goibibo.com/api/search/?app_id=cc7a88d8&app_key=a545d59e587dfc2fe1ba8dddf39b7b24&format=json&source=BLR&destination=IXR&dateofdeparture=20180614&dateofarrival=20180624&seatingclass=E&adults=1&children=0&infants=0&counter=100"
	resp, urlError := http.Get(url)

	if urlError != nil {
		fmt.Println(urlError)
		return nil, urlError
	}

	data, bodyError := ioutil.ReadAll(resp.Body)

	if bodyError != nil {
		fmt.Println(bodyError)
		return nil, urlError
	}
	var air models.AirfareResponse
	unmarshalError := json.Unmarshal(data, &air)

	if unmarshalError != nil {
		fmt.Println(unmarshalError)
		return nil, unmarshalError
	}

	return &air, nil

}

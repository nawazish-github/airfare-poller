package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/nawazish-github/airfare-poller/models"
)

type ConfigUnmarshaller struct{}

func (configUnmarshaller ConfigUnmarshaller) Unmarshal(file string) (*models.Config, error) {

	data, fileReadError := ioutil.ReadFile(file)
	if fileReadError != nil {
		fmt.Println(fileReadError)
		return nil, fileReadError
	}

	var config models.Config

	jsonUnmarshalError := json.Unmarshal(data, &config)
	if jsonUnmarshalError != nil {
		fmt.Println(jsonUnmarshalError)
		return nil, jsonUnmarshalError
	}

	return &config, nil

}

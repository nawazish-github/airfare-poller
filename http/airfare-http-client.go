package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type AirfareHttpClient struct{}

func (airfareHttpClient AirfareHttpClient) Get(url string) ([]byte, error) {
	//fmt.Println("GETting...", url)
	resp, urlError := http.Get(url)

	if urlError != nil {
		fmt.Println(urlError)
		return nil, urlError
	}

	data, bodyError := ioutil.ReadAll(resp.Body)

	if bodyError != nil {
		fmt.Println(bodyError)
		return nil, bodyError
	}

	//fmt.Println(data)
	return data, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	Origin string `json:"origin"`
}

func getApi() {
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(data, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %s\n", err)
		return
	}

	fmt.Printf("IP Address: %s\n", apiResponse.Origin)
}

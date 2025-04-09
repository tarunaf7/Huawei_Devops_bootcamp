













package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


var (
	packagingServiceUrl string
)

type PackagingInfo struct {
	Weight float32 `json:"weight"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Depth  float32 `json:"depth"`
}


func init() {
	packagingServiceUrl = os.Getenv("PACKAGING_SERVICE_URL")
}

func isPackagingServiceConfigured() bool {
	return packagingServiceUrl != ""
}

func httpGetPackagingInfo(productId string) (*PackagingInfo, error) {
	// Make the GET request
	url := packagingServiceUrl + "/" + productId
	fmt.Println("Requesting packaging info from URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	// Read the JSON response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decode the JSON response into a PackagingInfo struct
	var packagingInfo PackagingInfo
	err = json.Unmarshal(responseBody, &packagingInfo)
	if err != nil {
		return nil, err
	}

	return &packagingInfo, nil
}

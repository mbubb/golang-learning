package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to parse the API response
type MetalPriceResponse struct {
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
	Unit      string             `json:"unit"`
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
}

func main() {
	url := "https://api.metalpriceapi.com/v1/latest?api_key=91bec96031604730144166ea4ae40641&base=USD&currencies=BTC,ETH,XAU,XAG"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var data MetalPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Check if API call was successful
	if !data.Success {
		log.Fatalf("API call failed")
	}

	// Print the prices
	fmt.Printf("Base Currency: %s\n", data.Base)
	for currency, rate := range data.Rates {
		fmt.Printf("%s: %.4f %s\n", currency, rate, data.Unit)
	}
}

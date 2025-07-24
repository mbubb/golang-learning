package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Asset represents an investment type like gold or crypto
type Asset struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Amount   float64 `json:"amount"`
	USDPrice float64 `json:"usd_price"`
}

// APIResponse is used to parse CoinGecko or mock API response
type APIResponse map[string]struct {
	USD float64 `json:"usd"`
}

func main() {
	portfolio := []Asset{
		{Name: "Bitcoin", Symbol: "bitcoin", Amount: 0.25},
		{Name: "Ethereum", Symbol: "ethereum", Amount: 2.0},
		{Name: "Gold", Symbol: "gold", Amount: 1.5},      // 1.5 oz
		{Name: "Silver", Symbol: "silver", Amount: 20.0}, // 20 oz
	}

	prices := fetchPrices([]string{"bitcoin", "ethereum", "gold", "silver"})

	total := 0.0
	fmt.Println("\n--- Portfolio Summary ---")
	for i, asset := range portfolio {
		price := prices[asset.Symbol]
		portfolio[i].USDPrice = price
		value := asset.Amount * price
		total += value
		fmt.Printf("%-8s: %5.2f @ $%.2f = $%.2f\n", asset.Name, asset.Amount, price, value)
	}
	fmt.Printf("\nTotal Portfolio Value: $%.2f\n", total)
}

func fetchPrices(symbols []string) map[string]float64 {
	ids := ""
	for i, s := range symbols {
		ids += s
		if i != len(symbols)-1 {
			ids += ","
		}
	}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", ids)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error fetching prices:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var data APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	prices := make(map[string]float64)
	for k, v := range data {
		prices[k] = v.USD
	}

	return prices
}

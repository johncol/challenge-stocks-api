package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/", getStockHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func getStockHandler(writter http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol != "" {
		stock := searchStockForSymbol(symbol)
		json.NewEncoder(writter).Encode(stock)
	}
}

type Stock struct {
	Symbol string `json:"symbol"`
	Date   string `json:"date"`
	Value  string `json:"value"`
}

type RemoteStock struct {
	Metadata StockMetadata         `json:"Meta Data"`
	Values   map[string]StockValue `json:"Time Series (5min)"`
}

type StockValue struct {
	Close string `json:"4. close"`
}

type StockMetadata struct {
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
}

func searchStockForSymbol(symbol string) Stock {
	remoteStock := requestStockFromApi(symbol)
	stock := mapRawStockToStock(remoteStock)
	return stock
}

func requestStockFromApi(symbol string) RemoteStock {
	resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&interval=5min&apikey=XHPGM4KDYKBFKA4Q&symbol=" + symbol)
	handleErrorIfAny(err)
	defer resp.Body.Close()

	var remoteStock RemoteStock
	err = json.NewDecoder(resp.Body).Decode(&remoteStock)
	handleErrorIfAny(err)

	return remoteStock
}

func handleErrorIfAny(err error) {
	if err != nil {
		panic(err)
	}
}

func mapRawStockToStock(stockData RemoteStock) Stock {
	fmt.Print(stockData.Values)
	return Stock{
		Value:  stockData.Values[stockData.Metadata.LastRefreshed].Close,
		Symbol: stockData.Metadata.Symbol,
		Date:   stockData.Metadata.LastRefreshed,
	}
}

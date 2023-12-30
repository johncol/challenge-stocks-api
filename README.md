# Stocks API

Dummy API written in Go to get stock values as part of the interview process for Twilio Colombia.

Run with

    go run stock-api.go

Then in your browser go to http://localhost:8080/?symbol={STOCK-SYMBOL} and get a JSON response like:

    {
      "symbol": "{STOCK-SYMBOL}",
      "date": "2018-09-28 15:55:00",
      "value": "86.2800"
    }

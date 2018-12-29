## The Unofficial Finviz API for Golang

**Introduction**

[Finviz.com](http://www.finviz.com) aims to make market information accessible and provides a lot of data in visual snapshots, allowing traders and investors to quickly find the stock, future or forex pair they are looking for. This project aims to provide an unoffical api to screen for stocks on finviz.

### Important information

Any quotes data displayed on finviz.com is delayed by 15 minutes for NASDAQ, and 20 minutes for NYSE and AMEX. This API should **NOT** be used for live trading, it's main purpuse is financial analysis, research and data scraping.

### Install the latest package

    go get github.com/shitbox/finviz

### Using Screen(..)

    package main

    import (
	      "github.com/shitbox/finviz"
	      "fmt"
	   )
    
    func main() {
	    filter := []string{"ta_averagetruerange_o0.5", "ta_sma20_sa50" }
	    k := finviz.Screen(filter)
	    fmt.Println(k.ToJson())
    }

    ------------ Output ------->>
	"Stocks": [{
		"No": "1",
		"Ticker": "ADIL",
		"Company": "Adial Pharmaceuticals, Inc.",
		"Sector": "Healthcare",
		"Industry": "Biotechnology",
		"Country": "USA",
		"MarketCap": "31.64M",
		"PriceToEarning": "-",
		"Price": "5.55",
		"PercentageChange": "-3.98%",
		"Volume": "563,162"
	}, {
		"No": "2",
		"Ticker": "AGQ",
		"Company": "ProShares Ultra Silver",
		"Sector": "Financial",
		"Industry": "Exchange Traded Fund",
		"Country": "USA",
		"MarketCap": "-",
		"PriceToEarning": "-",
		"Price": "26.08",
		"PercentageChange": "1.52%",
		"Volume": "103,887"
	}, {
		"No": "3",
		"Ticker": "AMRC",
		"Company": "Ameresco, Inc.",
		"Sector": "Services",
		"Industry": "Technical Services"
    .
		.
		.
		.
		.
		.
		.
		.
		.
		.
		
		"No": "20",
		"Ticker": "LTPZ",
		"Company": "PIMCO 15+ Year US TIPS ETF",
		"Sector": "Financial",
		"Industry": "Exchange Traded Fund",
		"Country": "USA",
		"MarketCap": "-",
		"PriceToEarning": "-",
		"Price": "62.43",
		"PercentageChange": "0.33%",
		"Volume": "22,485"
	}],
	"Time": "2018-12-28T21:40:45.206318-08:00"
  }

### Documentation

More documentation upcoming with more features

### Contributing 

You can contribute to the project by reporting bugs, suggesting enhancements, or directly by extending and writing features.

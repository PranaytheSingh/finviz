package finviz

import (
	"testing"
)

func TestArrayToString(t *testing.T) {
	filtersString := "ta_averagetruerange_o0.5,ta_candlestick_lls,ta_sma20_sa50"
	var filter = []string{"ta_averagetruerange_o0.5", "ta_candlestick_lls", "ta_sma20_sa50"}
	k := arrayToString(filter)
	if k != filtersString {
		t.Error("Expected " + filtersString + "got " + k)
	}
}

func TestScrape(t *testing.T) {
	filtersString := "ta_averagetruerange_o0.5,ta_sma20_sa50"
	k := Scrape(filtersString)
	if k[0].No != "1" {
		t.Error("Did not recieve data array")
	}
}

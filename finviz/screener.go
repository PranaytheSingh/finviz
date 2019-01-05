package finviz
	
import (
	 "github.com/PuerkitoBio/goquery"
	 "log"
	 "fmt"
	 "encoding/json"
	 "time"
	 "strconv"
	 )

const (
site = "http://www.finviz.com/screener.ashx"
columns = 11
)

type response struct {
	Stocks []template
	Time time.Time
}

func (r response) ToJson() string {
	jsonData, err := json.Marshal(r)
    if err != nil {
        fmt.Printf("Error: %s", err)
    }	
    return string(jsonData)
}

type template struct {
	No string
	Ticker string
	Company string
	Sector string
	Industry string
	Country string
	MarketCap string
	PriceToEarning string
	Price string
	PercentageChange string
	Volume string
	Date string
}

func (t template) init(a []string) template {
	date := get_date()
	return template{a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], date}
}

func (t template) print() {
	fmt.Println("********" + t.Ticker + "********")
}

func get_date() string{
	year, month, day := time.Now().Date()
	date := strconv.Itoa(int(month))+"-"+strconv.Itoa(day)+"-"+strconv.Itoa(year)
	return date
}

func Screen(filters []string) response {
	filtersString := arrayToString(filters)
	data := Scrape(filtersString)
	t := time.Now()
	return response{data, t}
}

func Scrape(filters string ) []template {
	
	var values []string
	var iterates int
	var dataarray []template
	var secondOccurance int

	bod := MakeRequest(site, filters)

	doc, err := goquery.NewDocumentFromReader(bod)
  	
  	if err != nil {
    	log.Fatal(err)
  	}

  	doc.Find("table").Each(func(i int, tablehtml *goquery.Selection) {
    	
    	tablehtml.Find("[valign=\"top\"]").Each(func(i int, tablerow *goquery.Selection) {
   			
   			tablerow.Find("td").Each(func(i int, td *goquery.Selection){
   	
   			values = append(values, td.Text())
    		
    		})
    	})
	})

  	if len(values) == 1 {
  		log.Fatal("No stocks to fetch")
  	}

  	values = values[1:] //Remove first garbage elemet

  	firstStock := values[1] //Get the first stock name

	for index, element := range values {
		if element == firstStock{
			secondOccurance = index // Get next occurance of first stock
		}
	}

	values = values[0:secondOccurance-1] // Split away all second occurances of the stocks

	iterates = len(values)/columns

	i := 0
	j := 0
	for i < iterates {
		
		temparr := values[j:11+(11*i)] //Get first 11 from list 
		data := template{}.init(temparr)
		dataarray = append(dataarray, data)

		i++
		j = j + columns
	}

	return dataarray
}
package finviz

import(
	"net/http"
    "log"
    "io"
)



func MakeRequest(url string, headers string) io.Reader {

    URL := url + "?f=" + headers
    req, _ := http.NewRequest("GET", URL, nil)
    log.Println("Requesting.."+URL)
    res, err := http.DefaultClient.Do(req)
    
    log.Println(res.StatusCode)
    
    if err != nil {
        log.Fatal(err)
    }

    peabody := res.Body
    return peabody
}

package finviz

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(url string, headers string) (io.Reader, error) {

	URL := url + "?f=" + headers
	req, _ := http.NewRequest("GET", URL, nil)
	fmt.Println("Requesting.." + URL)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res.StatusCode)

	peabody := res.Body
	return peabody, nil
}

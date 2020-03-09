package httputils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequestBody(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[%d] %s\n", res.StatusCode, res.Request.URL)
	return body
}

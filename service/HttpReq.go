package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpReqGet(url string) []byte {

	// set timeout
	TimeOut := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	// set req to get data
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set headers
	req.Header.Set("User-Agent", "GITS-INDONESIA")

	// get response api
	res, getErr := TimeOut.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	//parsing body
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body

}

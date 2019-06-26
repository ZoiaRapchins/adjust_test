package main

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"net/http"
)

type displayResponseType func(string, [16]byte)

func DoGetRequest(url string, finished chan<- bool, displayResponse displayResponseType) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		finished <- true
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		finished <- true
		return
	}

	displayResponse(url, md5.Sum(body))
	finished <- true
}

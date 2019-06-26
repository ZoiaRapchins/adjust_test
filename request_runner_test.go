package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"testing"
)

func runMockedServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test")
	})
	http.ListenAndServe(":80", nil)
}

const TestUrl = "http://localhost:80"
const TestChecksum = "0cbc6611f5540bd0809a388dc95a615b"

func TestDoGetRequest(t *testing.T) {
	go runMockedServer()

	displayFunction := func(url string, md5Response [16]byte) {
		if url != TestUrl {
			t.Fatalf("Wrong url (expected: %s, actual: %s)", TestUrl, url)
		}

		if hex.EncodeToString(md5Response[:]) != TestChecksum {
			t.Fatalf("Wrong checksum (expected: %x, actual: %s)", TestChecksum, md5Response)
		}
	}

	finished := make(chan bool, 1)

	DoGetRequest("http://localhost:80", finished, displayFunction)

	<-finished
}

func TestDoGetRequest_InvalidUrl(t *testing.T) {
	displayIsNotTriggered := true
	displayFunction := func(url string, md5Response [16]byte) {
		displayIsNotTriggered = false
	}

	finished := make(chan bool, 1)

	DoGetRequest("wrong", finished, displayFunction)

	<-finished

	if !displayIsNotTriggered {
		t.Fatalf("DisplayFunction was triggered but was not expected")
	}
}

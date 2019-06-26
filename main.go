package main

import (
	"fmt"
	"os"
)

func main() {
	cmdArgs := ParseArguments(os.Args[1:])

	displayFunction := func(url string, md5Response [16]byte) {
		fmt.Printf("%s %x\n", url, md5Response)
	}

	executeGetRequest := func(url string, finished chan<- bool) {
		DoGetRequest(url, finished, displayFunction)
	}

	ExecuteRequests(cmdArgs, executeGetRequest)
}

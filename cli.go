package main

import (
	"flag"
	"strings"
)

const httpPrefix = "http://"
const defaultParallelRequests = 10

type CmdArguments struct {
	Urls                    []string
	NumberOfParallelRequest int
}

func ParseArguments(args []string) CmdArguments {
	requestsNumber := flag.Int("parallel", defaultParallelRequests, "maximum number of requests to run parallel")

	flag.CommandLine.Parse(args)

	parsedArgs := flag.Args()

	sanitizeUrls(parsedArgs)

	return CmdArguments{parsedArgs, *requestsNumber}
}

func sanitizeUrls(urls []string) {
	for index, url := range urls {
		if !strings.HasPrefix(url, httpPrefix) {
			urls[index] = httpPrefix + url
		}
	}
}

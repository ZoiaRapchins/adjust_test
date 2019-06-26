package main

import "testing"

func runTest(t *testing.T, maxParallelExecutions int) {
	testUrls := []string{"http://test1.com", "http://test2.com",
		"http://test3.com", "http://test4.com", "http://test5.com"}
	cmdArgs := CmdArguments{testUrls, maxParallelExecutions}

	numberOfTestUrls := len(testUrls)

	mockExecuteRequestFunc := func(url string, finished chan<- bool) {
		numberOfTestUrls = numberOfTestUrls - 1
		finished <- true
	}

	ExecuteRequests(cmdArgs, mockExecuteRequestFunc)

	if numberOfTestUrls != 0 {
		t.Error("Failed to execute all requests")
	}
}

var testMaxParallelExecutions = []int{10, 3}

func TestExecution(t *testing.T) {
	for _, maxParallelExecutions := range testMaxParallelExecutions {
		runTest(t, maxParallelExecutions)
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_SuccessfulParsedAndSanitized(t *testing.T) {
	testCmdArgs := []string{"-parallel=3", "http://test.com", "test2.com"}
	expectedUrls := []string{"http://test.com", "http://test2.com"}

	result := ParseArguments(testCmdArgs)

	if result.NumberOfParallelRequest != 3 {
		t.Error("Incorrect NumberOfParallelRequest")
	}

	if !reflect.DeepEqual(result.Urls, expectedUrls) {
		t.Error("Failed to parse and sanitize urls")
	}
}

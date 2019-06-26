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

	ExecuteRequests(cmdArgs, displayFunction)
}

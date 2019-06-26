package main

type ExecuteRequestType func(url string, finished chan<- bool)

func ExecuteRequests(cmdArgs CmdArguments, executeRequestType ExecuteRequestType) {
	numberOfRunningJobs := cmdArgs.NumberOfParallelRequest
	finished := make(chan bool, numberOfRunningJobs)
	for i := 0; i < numberOfRunningJobs; i++ {
		if len(cmdArgs.Urls) == 0 {
			numberOfRunningJobs = i
			break
		}
		url := cmdArgs.Urls[len(cmdArgs.Urls)-1]
		cmdArgs.Urls = cmdArgs.Urls[:len(cmdArgs.Urls)-1]
		go executeRequestType(url, finished)
	}

	for i := 0; i < numberOfRunningJobs; i++ {
		<-finished
		if numberOfRunningJobs == 0 {
			break
		}

		if len(cmdArgs.Urls) > 0 {
			numberOfRunningJobs = numberOfRunningJobs + 1
			url := cmdArgs.Urls[len(cmdArgs.Urls)-1]
			cmdArgs.Urls = cmdArgs.Urls[:len(cmdArgs.Urls)-1]
			go executeRequestType(url, finished)
		}
	}
}

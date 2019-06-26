package main

func ExecuteRequests(cmdArgs CmdArguments, displayFunction displayResponseType) {
	numberOfRunningJobs := cmdArgs.NumberOfParallelRequest
	finished := make(chan bool, numberOfRunningJobs)
	for i := 0; i < numberOfRunningJobs; i++ {
		if len(cmdArgs.Urls) == 0 {
			numberOfRunningJobs = i
			break
		}
		url := cmdArgs.Urls[len(cmdArgs.Urls)-1]
		cmdArgs.Urls = cmdArgs.Urls[:len(cmdArgs.Urls)-1]
		go DoGetRequest(url, finished, displayFunction)
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
			go DoGetRequest(url, finished, displayFunction)
		}
	}
}

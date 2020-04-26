package main

import "fmt"


func main() {
	fmt.Println("start the function ")
	port := 3000
	_, errorStart := StartWebService(port, 2)  // Because declared variable must be used, so use _ to avoid use it and can be approved by complier 
	fmt.Println(errorStart)
}

// StartWebService is about how to start the web service
func StartWebService(port int, retryOfTimes int) (int, error) {
	fmt.Println("Start the web service ... ")
	// Do something
	fmt.Println("Web Service has been listened on port : ", port)
	fmt.Println(retryOfTimes, "has been retried")
	return port, nil
}
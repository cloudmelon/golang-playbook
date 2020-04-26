package main

import "fmt"

func main() {
	fmt.Println("start the function ")
	port := 3000
	StartWebService(port, 2)
}

// StartWebService is about how to start the web service
func StartWebService(port int, retryOfTimes int) {
	fmt.Println("Start the web service ... ")
	// Do something
	fmt.Println("Web Service has been listened on port : ", port)
	fmt.Println(retryOfTimes, "has been retried")

	fmt.Println("Start the web service successfully ")

}

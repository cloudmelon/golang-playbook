package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("start the function ")
	port := 3000
	errorStart := StartWebService(port, 2)
	fmt.Println(errorStart)
}

// StartWebService is about how to start the web service
func StartWebService(port int, retryOfTimes int) error {
	fmt.Println("Start the web service ... ")
	// Do something
	fmt.Println("Web Service has been listened on port : ", port)
	fmt.Println(retryOfTimes, "has been retried")
	return errors.New("there is an error")
}

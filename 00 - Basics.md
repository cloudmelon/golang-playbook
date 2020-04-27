# Play 0 - Getting Started

This page introduces basics of Golang as a start point. 

## Installing the Go Tools

You can download the go package at the website and follow the guidance to install the tool: https://golang.org/dl/

After install the tools, go and check the version : 

    ```
      go version check the current version 
    
    ```
    
Install all dependencies :

    ```
      go install // install and complie packages and dependencies
    ```

To find all go packages : https://golang.org/pkg/

## Input and print out variables 

When you're ready, just show you some quick examples : 


    ```
    // this the first comment 
    package main

    import (
        "fmt"
    )

    /*
    this is a mult-line comment 
    
    */
  

    func main() {
        fmt.Println("Hello world ")
    }

    ```

## Go application by modules

If you decide to go for Go modules, you can start your Go project by using **go mod init** command, such as the following command if you wanna create a module named **cloudmelon/helloworldmodule** :

 ```
    go mod init cloudmelon/hellowordmodule

 ```

Expecting the output would be as the following : 
<img src="Screenshots/go module.PNG" alt="go module" width="600px"/>

If you go to go.mod file you would see something similar to the following : 
    ```
    module cloudmelon/hellowordmodule

    go 1.14
    ```

## Run Go application

After you're done with your coding, you can use the following command to run the module which calls your **Main** function 

 ```
    go run cloudmelon/hellowordmodule

 ```

Otherwise, if you're already in the folder contains **go.mod** file, you can use **period** instead

 ```
    go run . 

 ```

Expecting the output would be as the following : 
<img src="Screenshots/run module.PNG" alt="run module" width="600px"/>


## Build Go application
If you're under window OS, using **go build** command will build up a windows executable by default. Similarily, you can use the following : 

```
    go build < full module name >
    
    go build . 

 ```


## Test Go application

Use **go test** command to test the application

    ```
    go test -v

    ```

If the test suite contains benchmarks, you can run these with the --bench and --benchmem flags


    ```
    go test -v --bench . --benchmem

    ```

More Unit test example on Go, check it out :  https://gobyexample.com/testing


## Panic function


    ```
    panic("something goes wrong")

    ```




## Best practices

- Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. 

https://golang.org/doc/effective_go.html

- Go project layout : 

https://github.com/golang-standards/project-layout

- This package implements an HTTP request pipeline suitable for use across multiple go-routines and provides the shared routines relied on by AutoRest generated Go code. https://github.com/Azure/autorest.go/

https://github.com/Azure/go-autorest


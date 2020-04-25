## Play 0 - Getting Started

This page introduces basics of Golang as a start point. 

## Installing the Go Tools

You can download the go package at the website and follow the guidance to install the tool: https://golang.org/dl/

After install the tools, go and check the version : 

    ```
      go version check the current version 
    
    ```


## Input and print out variables 

When you're ready, just show you some quick examples : 


    ```
   package main

    // this the first comment 

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

## Common GO CLI comands 

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


After you're done with your coding, you can use the following command to run the module which calls your **Main** function 

 ```
    go mod run cloudmelon/hellowordmodule

 ```

Expecting the output would be as the following : 
<img src="Screenshots/run module.PNG" alt="run module" width="600px"/>




Utimately, the following commands are quite useful :

    ```
      go install // install and complie packages and dependencies
     
      go build  // build a go program

      go test // test go program packages

      go run // Compile + run Go program

    ```
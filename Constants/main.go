package main

import "fmt"

// define constant  
const constantName = "hello constant"

// constant blocks ( evaluation at compile time not runtime )
const (
	first = 1
	second = iota + 1 
)

// iota is gonna reset each time 
const (
	third = iota
	fourth = iota
)
	
func main() {

	fmt.Println(constantName)

	fmt.Println(second)

	fmt.Println(first, second, third, fourth)
}

	/*

		Output :
		hello constant
		2
		1 2 0 1
		
	*/
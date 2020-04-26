package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 5, 7, 8, 9)
	fmt.Println(slice)

	s2 := slice[1:2]
	s3 := slice[1:]
	s4 := slice[:2]
	fmt.Println(s2, s3, s4)

}

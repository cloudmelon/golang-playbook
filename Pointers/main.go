package main

import "fmt"

func main() {

	/*	var firstName *string = new(string)
	*firstName = "Melon" //dereference the value
	fmt.Println(*firstName)*/

	firstName := "Cloudmelon"

	ptr := &firstName //
	fmt.Println(ptr, *ptr)

	firstName = "MSmelon"
	fmt.Println(ptr, *ptr)

	/*

			Output :
			0xc0000381f0 Cloudmelon
		    0xc0000381f0 MSmelon

			The value that pointer was pointed to changed, whereas memory address doesn't change ( data stored there has been changed )

	*/
}

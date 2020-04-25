package main

import "fmt"

func main() {

	// complex type
	com := complex(3, 4) // In this case, it means (3+4i)
	fmt.Println(com)
	r, i := real(com), imag(com)
	fmt.Println(r, i)

	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"
}

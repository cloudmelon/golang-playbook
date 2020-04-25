package main

import "fmt"

func main() {

	// Declarate the data type
	var i int
	i = 1
	fmt.Println(i)

	var f32 float32 = 3.14
	var f64 float64 = 3.1415
	fmt.Println(f32)
	fmt.Println(f64)

	// implicit initialization
	// string
	stringName := "cloudmelon"
	fmt.Println(stringName)

	// float32 or float64
	floatNumber := 3.14
	fmt.Println(floatNumber)

	// boolean ( use bool reserved word to define it )
	boolean := true
	fmt.Println(boolean)

	// complex type
	com := complex(3, 4) // In this case, it means (3+4i)
	fmt.Println(com)
}

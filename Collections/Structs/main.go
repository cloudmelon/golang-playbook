package main

import (
	"fmt"
)

func main() {
	type cat struct {
		name string
		age  int
	}

	var c cat
	c.name = "Lucie"
	c.age = 1
	fmt.Println(c)

	c2 := cat{
		name: "Acer",
		age:  2,
	}
	fmt.Println(c2)

}

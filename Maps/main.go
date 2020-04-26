package main

import (
	"fmt"
)

func main() {
	m :=map[string]int{"cat":5}
	fmt.Println(m)
	fmt.Println(m["cat"])

	m["cat"] = 6
	fmt.Println(m)

	delete(m,"cat")
	fmt.Println(m)

}

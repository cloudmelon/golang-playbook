package main

func main() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

}

/*

	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

*/

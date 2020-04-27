package main

func main() {
	ports := map[string]int{"http": 80, "https": 443}
	for i, v := range ports {
		println(i, v)
	}

}

/*

ports := map[string]int{"http": 80, "https": 443}
	for _, v := range ports {
		println(v)
	}


ports := map[string]int{"http": 80, "https": 443}
	for i := range ports {
		println(i)
	}

*/

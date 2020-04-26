# Play 2 - Collections


## Arrays

Fixed size array. 

    ```
    package main

    import (
        "fmt"
    )

    func main() {
        var arr [3]int
        arr[0] = 1
        arr[1] = 2
        arr[2] = 3

        fmt.Println(arr)

        arr2 := [3]int{1, 2, 3}
        fmt.Println(arr2)
    }

    /*
		Output :
		[1 2 3]
        [1 2 3]
	*/
    ```

## Slices 

Based on the array, however can be resize as needed. 

    ```
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

     /*
		Output :
		[1 2 3]
        [1 2 3 5 7 8 9]
        [2] [2 3 5 7 8 9] [1 2]
	*/
    
    ```

## Maps 

Map is like a key value dictionary 

   ```
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

     /*
		Output :
		map[cat:5]
        5
        map[cat:6]
        map[]
	*/
    
    ```

## Structs

You can define a struct at any level of your application :

```
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

    /*
		Output :
		{Lucie 1}
        {Acer 2}
	*/

```
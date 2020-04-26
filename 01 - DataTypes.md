## Play 1 - DataTypes 

This page introduces basics of data types of Go. 


## Declaire data types

    ```
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

    ```

Expecting the output would be as the following : 
<img src="Screenshots/data type.PNG" alt="data type" width="600px"/>


## Working with Constants

You can simply use **const** keyword to define a constant type. it is not supposed to change constant value in the program, however it could work if you remember to conversion at the time. 

    ```
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
    ```




## Complex Numbers

Go provides two sizes of complex numbers, **complex64** and **complex128**, whose components are **float32** and **float64** respectively. The built-in function complex creates a complex number from its **real** and **imaginary** components, and the built-in real and imag functions extract those components:



    ```
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

    ```

## Working with Pointers

Instead of hold the value of a variable directly, we're going to use pointer to hold up the address of a location in memory. You'll need to initialize the pointer firstly and then use asterix to deference the value as the follows : 

    ```
    package main

    import "fmt"

    func main() {

        var firstName *string = new(string)
        *firstName = "Melon" //dereference the value
        fmt.Println(*firstName)

    }

    ```

The other way : 

    ```
    package main

    import "fmt"

    func main() {

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

    ```

Using pointers to collaborate with data without have the copy of data around across different parts of application. 


## Refs 

https://www.interviewcake.com/concept/java/bit-shift
package main

import "fmt"

func main() {

	// Complex Numbers

	// way-1
	var a complex64 = 5 + 5i
	var b complex64 = 10 + 10i

	c := a + b

	fmt.Println(c)

	// way-2
	var z complex64 = complex(3, 4)
	// var y = complex(2, 1)
	var y complex64 = complex(2, 1)

	fmt.Println(z)
	fmt.Println(y)

	var d complex64 = z + y

	fmt.Println(d)

	// Printing  the type
	fmt.Printf("The type of d is %T", d)

}

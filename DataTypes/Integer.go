package main

import "fmt"

func main() {

	// Signed Integers
	var score int = 500 // int - Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems
	var balance int = -100

	var a int8 = -50 // 8bits range

	// Unsigned Integers
	var id uint = 1000 // unit - Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems

	var b uint16 = 40

	fmt.Println(score)
	fmt.Println(balance)
	fmt.Print(a)
	fmt.Println(id)
	fmt.Println(b)

}

//  Refer For diff range - https://www.geeksforgeeks.org/data-types-in-go/

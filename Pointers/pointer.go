// Pointer or Special variable that stores the memory address of another variable.

package main

import "fmt"

func main() {
	var name string
	var namePointer *string

	namePointer = &name
	*namePointer = "Lakshay"

	fmt.Println(name)
	fmt.Println(&name)
	println(namePointer)
	fmt.Printf("The name of user stored at %v location is %v \n", namePointer, *namePointer)
	fmt.Printf("The name of user stored at %v location is %v \n", namePointer, name)
}

package main

import "fmt"

func main() {
	// Note:- In Go Arrays not having elements same as the length will print space or 000 as below.
	var customerNames = [50]string{"Alex", "Le", "Hesh"}

	// idNo := [100]int{01, 02, 03, 04}
	idNo := [4]int{01, 02, 03, 04}

	fmt.Println(customerNames)
	fmt.Println(idNo)

	fmt.Printf("The customer named %v belongs to id no %v.", customerNames[0], idNo[0])
}

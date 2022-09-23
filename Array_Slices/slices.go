package main

import "fmt"

func main() {
	// Note:- In Go Arrays not having elements same as the length will print space or 000 as below.
	var customerNames = []string{"Alex", "Le", "Hesh"} //slice

	// idNo := [100]int{01, 02, 03, 04}
	idNo := [4]int{01, 02, 03, 04} //array

	fmt.Println(customerNames)
	fmt.Println(idNo)

	// customerNames[5] = "Yun"
	customerNames = append(customerNames, "Yun") // Adds the element at the end of the slice. Now we don't need to take track of indexes for addinf elements.
	//Grows the size if a greater capacity is needed and returns the updated slice value.

	fmt.Println(customerNames)

	fmt.Printf("The customer named %v belongs to id no %v.", customerNames[0], idNo[0])
}

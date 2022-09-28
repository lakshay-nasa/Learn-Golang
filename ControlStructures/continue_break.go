package main

import "fmt"

func main() {
	var name = []string{"pop", "winny", "stranger", "yin", "luke", "tian"}

	// The Continue Statement
	fmt.Println("\nThe Continue Statement")
	for i := 0; i < len(name); i++ {
		if name[i] == "stranger" {
			continue
		}

		fmt.Printf("Welocme! %v \n", name[i])
	}

	// The Break Statement
	fmt.Println("\nThe Break Statement")
	for i := 0; i < len(name); i++ {
		if name[i] == "stranger" {
			fmt.Println("Warning!!")
			break
		}

		fmt.Printf("Welocme! %v \n", name[i])
	}
}

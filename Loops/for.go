package main

import "fmt"

func main() {

	// Infinte Loop
	// for {
	// 	fmt.Println("Hello")
	// }

	// Counter Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

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

package main

import "fmt"

func main() {
	// Go Conditional Statements

	// The if Statement

	x := 4
	y := 9

	if x < y {
		fmt.Printf("%v is less than %v \n", x, y)
	}

	if x > y {
		fmt.Printf("%v is greater than %v \n", x, y)
	}

	// The if-else Statement

	if x < y {
		fmt.Printf("%v is less than %v \n", x, y)
	} else {
		fmt.Printf("%v is greater than %v \n", x, y)
	}

	// The if-else-if Statement
	time := 20
	if time < 12 {
		fmt.Println("Good Morning")
	} else if time < 18 {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good Evening")
	}

	// Nested if Statement
	if x < y {
		if x == 4 {
			fmt.Printf("%v is equal to 4 \n", x)
		}
	}
}

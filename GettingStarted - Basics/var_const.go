package main

import "fmt"

func main() {
	const PI = 3.14
	var temp = "280 K"

	fmt.Print("The value of PI is", PI, " which is constant.")
	fmt.Print("The value of temprature is", temp, " which is can vary.\n")

	fmt.Println("The value of PI is", PI, " which is constant.")
	fmt.Println("The value of temprature is", temp, " which is can vary.")

	fmt.Printf("The value of PI is %v which is constant.", PI)
	fmt.Printf("The value of temprature is %v which is can vary.\n", temp)
}

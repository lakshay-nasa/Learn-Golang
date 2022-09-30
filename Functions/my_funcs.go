package main

import "fmt"

func myfunction() {
	fmt.Println("Hello, Stranger!")
}

func userInput() (string, int, string) {
	var name, email string
	var age int
	println("What is your name?")
	fmt.Scan((&name))
	println("What is your age?")
	fmt.Scan((&age))
	println("Enter your email address:")
	fmt.Scan((&email))
	return name, age, email
}

func userDetails() {
	name, age, email := userInput()

	fmt.Printf("Hello %s, you are %d years old. Your email address is %s \n", name, age, email)
}

func main() {
	myfunction()
	userDetails()
}

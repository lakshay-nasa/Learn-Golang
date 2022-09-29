package main

import "fmt"

func main() {
	var name string
	println("Enter your name: ")
	fmt.Scan(&name)

	switch name {
	case "yun":
		fmt.Printf("Hello, %s! you are a member.", name)
	case "rin":
		fmt.Printf("Hello, %s! you are a member.", name)
	case "jinny":
		fmt.Printf("Hello, %s! you are a member.", name)
	default:
		fmt.Printf("Sorry, %s you are not a member.", name)
	}
}

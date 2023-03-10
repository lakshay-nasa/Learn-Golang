package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Greeting string `json:"greeting"`
}

var users []User

func main() {
	router := gin.Default()
	router.Run("localhost:8080")
}

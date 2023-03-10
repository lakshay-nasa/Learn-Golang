package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
)

func main() {

	resp, err := http.Get("API_Link")

	if err != nil {
		fmt.Print(err.Error())
		// os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respData))
}

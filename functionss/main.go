package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Inside main function")
	port := 3000
	webServerStatus := startWebserver(port)
	fmt.Println(webServerStatus)
	errval, err := errorLogCheck()
	fmt.Println(errval, err)
}

func startWebserver(port int) int {
	fmt.Println("Web server starting at port ",port)
	fmt.Println("Web server started.")
	return 200
}

func errorLogCheck() (int, error) {
	fmt.Println("any errors??")
	return 100, errors.New("something wemt wrong")
}
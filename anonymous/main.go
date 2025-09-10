package main

import "fmt"

func main() {
	a := func(name string) {
		fmt.Println("Hi", name)
		//return name
	}

	a("jerome")
	a("john")
}
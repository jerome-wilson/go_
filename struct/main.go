package main

import (
	"fmt"
)
type user struct {
	ID int
	name string
	age int
}

func main() {

	var u user
	u.ID = 1
	u.name = "Jerome"
	u.age = 22

	fmt.Println("Method 1")
	fmt.Println(u)
	fmt.Println(u.ID)
	fmt.Println(u.name)
	fmt.Println(u.age)

	fmt.Println()

	fmt.Println("Method 2")
	u2 := user {
		ID: 1, 
		name: "Jerome", 
		age: 22,
	}
	fmt.Println(u2)
}
package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "Aurthur"
	fmt.Println(firstName)
	fmt.Println(*firstName)
}

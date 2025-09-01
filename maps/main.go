package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"idNum": 1, "age": 22}
	fmt.Println("Map: ", m)
	fmt.Println(m["idNum"])
	m["idNum"] = 2
	fmt.Println(m["idNum"])
	fmt.Println(m["age"])
	delete(m, "age")
}

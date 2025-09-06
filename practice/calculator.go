package main

import "fmt"

func takeValue() {
	var a, b int

	fmt.Print("Enter the value for a: ")
	fmt.Scan(&a)
	fmt.Println()
	fmt.Print("Enter the value for b: ")
	fmt.Scan(&b)
	calculator(a, b)
}

func calculator(a int, b int) {
	add := a + b
	subtract := a - b
	multiply := a * b
	divide := a /b
	modulo := a %b

	fmt.Println(a," + ",b," = ",add)
	fmt.Println(a," - ",b," = ",subtract)
	fmt.Println(a," * ",b," = ",multiply)
	fmt.Println(a," / ",b," = ",divide)
	fmt.Println(a," % ",b," = ",modulo)
}


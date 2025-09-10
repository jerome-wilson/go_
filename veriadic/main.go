package main

import "fmt"

func main() {
	fmt.Println(sum(2,43,6,56,7,6))
}

func sum(values ...float64) float64 {
	result := 0.0
	for _, value := range values {
		result += value
	}

	return result
}
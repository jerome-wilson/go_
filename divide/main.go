package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := divide(1, 2)

	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
	} else {
		fmt.Println((result))
	}
	
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return math.NaN(), errors.New("cannot divide by 0")
	}

	return a/b, nil
}
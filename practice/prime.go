package main

import (
	"errors"
	"fmt"
)

func isPrime(n int32) (bool, error) {
	if n < 0 {
		return false, errors.New("number smaller than zero")
	}
	if n < 2 {
		return false, errors.New("number should not be under 2")
	}

	for i := int32(2); i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func prime() {
	var number int32
	fmt.Print("Enter the number: ")
	fmt.Scan(&number)

	result, err := isPrime(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if result {
		fmt.Println("Yes!! It is a prime number")
	} else {
		fmt.Println("No, it is not a prime number")
	}
}

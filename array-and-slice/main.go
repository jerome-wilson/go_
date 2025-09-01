package main

import(
	"fmt"
)

func main() {
	var arr [3]int
	arr1 := [6]int{1,2,3,4,5,6}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Print("Array Declaration method 1: ")
	fmt.Println(arr)
	fmt.Print(arr[0])
	fmt.Print(arr[1])
	fmt.Println(arr[2])

	fmt.Print("Array Declaration method 2: ")
	fmt.Println(arr1)
	fmt.Println(arr1[3])

	fmt.Print("Using slice: ")
	slice := []int{1,2,3}
	fmt.Println(slice)
	slice = append(slice, 4, 5, 6)
	fmt.Println("After appending: ")
	fmt.Println(slice)
	
	s1 := slice[1:]
	s2:= slice[:2]
	s3 := slice[1:2]
	fmt.Println("Sliced array ", s1, s2, s3)
}
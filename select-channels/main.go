package main

import (
	"fmt"
	"time"
	//"sync"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	time.Sleep(10 * time.Millisecond)

	select {
	case mssg := <- ch1:
		fmt.Println(mssg)
	case mssg := <- ch2:
		fmt.Println(mssg)
	default:
		fmt.Println("No channels available")
	}
}
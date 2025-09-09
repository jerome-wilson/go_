package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter what you want to scream")
	n := bufio.NewReader(os.Stdin)
	s, _ := n.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s+"!!")
}
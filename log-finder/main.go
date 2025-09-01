package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
	"flag"
)


func main() {

	path := flag.String("path", "app.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are INFO, DEBUG, ERROR, and CRITICAL")
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
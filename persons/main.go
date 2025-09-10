package main

import (
	"fmt"
	//"errors"
)
func main() {
	p := NewPerson("Jerome", "Wilson")

	println(p.ID())
	println(p.FullName())
	err := p.SetTwitterHandler("@jerome")
	if err != nil {
		fmt.Println("An error occured during setting twiiter handler %s\n",err.Error())
	}
	println(p.TwiterHandler())
}

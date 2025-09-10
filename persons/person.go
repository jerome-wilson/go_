package main

import (
	"fmt"
	"strings"
	"errors"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastName  string
	twitterHandler string
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *Person) ID() string {
	return "1234"
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%v %v", p.firstName, p.lastName)
}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("Twitter handler must start with @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p Person) TwiterHandler() string {
	return p.twitterHandler
}
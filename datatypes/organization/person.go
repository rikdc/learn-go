package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName     string
	lastName      string
	twitterHandle string
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p Person) SetTwitterHandle(handle string) error {
	if len(handle) == 0 {
		p.twitterHandle = handle
	} else if !strings.HasPrefix(handle, "@") {
		return errors.New("twittter handle must start with an @ symbol")
	}
	p.twitterHandle = handle
	return nil
}

func (p Person) ID() string {
	return "12345"
}

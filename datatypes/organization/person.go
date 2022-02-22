package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandle string

func (th TwitterHandle) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandle TwitterHandle
	identifiable  Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		identifiable: identifiable,
	}
}

func (p *Person) ID() string {
	return p.identifiable.ID()
}

func (p Person) TwitterHandle() TwitterHandle {
	return p.twitterHandle
}

func (p *Person) SetTwitterHandle(handle TwitterHandle) error {
	if len(handle) == 0 {
		p.twitterHandle = handle
	} else if !strings.HasPrefix(string(handle), "@") {
		return errors.New("twittter handle must start with an @ symbol")
	}
	p.twitterHandle = handle
	return nil
}

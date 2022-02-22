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
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
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

func (p Person) ID() string {
	return p.FullName()
}

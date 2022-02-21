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

type Person struct {
	firstName     string
	lastName      string
	twitterHandle TwitterHandle
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
	return "12345"
}

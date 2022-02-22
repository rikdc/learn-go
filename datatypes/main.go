package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Tom", "Smith", organization.NewSocialSecurityNumber("123-45-6789"))
	err := p.SetTwitterHandle("@tom_smith")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err.Error())
	}
	println(p.ID())
	println(p.FullName())
	println(p.TwitterHandle())
	println(p.TwitterHandle().RedirectUrl())
}

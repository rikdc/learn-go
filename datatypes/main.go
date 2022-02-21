package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Tom", "Smith")
	err := p.SetTwitterHandle("@tom_smith")
	if err != nil {
		fmt.Printf("An error ocurred: %s", err.Error())
	}
	println(p.FullName())
	println(p.TwitterHandle())
	println(p.TwitterHandle().RedirectUrl())
}

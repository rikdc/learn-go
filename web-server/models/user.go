package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // A slice of pointers to user objects.
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {

	if u.ID == 0 {
		// We have to return a NullObject of User since
		// this function expects a return value.
		return User{}, errors.New("New user must include an id")
	}

	u.ID = nextId
	nextId++
	users = append(users, &u)

	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, k := range users {
		if k.ID == id {
			// remember to dereference the user object.
			return *k, nil
		}
	}
	return User{}, fmt.Errorf("User '%v' not found", id)
}

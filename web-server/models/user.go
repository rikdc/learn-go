package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // A slice of pointers to user objects.
	nextId = 1
)

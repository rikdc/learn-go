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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextId
	nextId++
	users = append(users, &u)

	return u, nil
}

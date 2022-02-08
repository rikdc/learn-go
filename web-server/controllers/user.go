package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from the User Controller!"))
}

// Constructor function.
// This uses a convention, not a rule.
func newUserController() *userController {
	// Returning the addressOf operator here is only applicable to structs.
	// This is a local variable, and returning the address.
	// C++ engineers may find this weird, but Go adds protections around this
	// by promoting the local variables scope.
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

package controllers

import "net/http"

type userController struct{}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from the User Controller!"))
}

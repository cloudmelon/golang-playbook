package controllers

import "net/http"

type userController struct {
}

// type wr Http information coming from that web request
func (uc userController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller "))
}

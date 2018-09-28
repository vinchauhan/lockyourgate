package controllers

import (
	"lock-my-gate/views"
	"net/http"
)

func NewUser() *User {
	return &User{
		NewView: views.NewView("semantic", "views/user/new.gohtml"),
	}
}

type User struct {
	NewView *views.View
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}
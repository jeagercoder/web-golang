package handler


import (
	"fmt"
	"net/http"
)


func UserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login")
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Register")
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Profile")
}
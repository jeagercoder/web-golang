package handler


import (
	"fmt"
	"net/http"
)


func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login")
}
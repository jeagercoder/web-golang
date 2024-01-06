package router


import (
	"net/http"
	"github.com/jeagercoder/web-golang/pkg/user/handler"
)


func SetupRouters() {
	http.HandleFunc("/login", handler.UserLogin)
	http.HandleFunc("/register", handler.UserRegister)
	http.HandleFunc("/profile", handler.UserProfile)
}
package middleware


import (
	"fmt"
	"net/http"
)


func LogMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)
}
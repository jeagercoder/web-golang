package main

import (
	"net/http"
	user_router "github.com/jeagercoder/web-golang/pkg/user/router"
)

func main()  {
	user_router.SetupRouters()
	http.ListenAndServe(":8000", nil)
}



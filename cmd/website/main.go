package main

import (
	"github.com/jeagercoder/web-golang/internal/app"
	"github.com/jeagercoder/web-golang/internal/middleware"
	user_router "github.com/jeagercoder/web-golang/pkg/user/router"
)

func main()  {
	app := app.New()
	app.Middleware(middleware.LogMiddleware)
	user_router.Setup(app)
	app.RunServer()
}



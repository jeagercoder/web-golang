package router


import (
	"github.com/jeagercoder/web-golang/internal/app"
	"github.com/jeagercoder/web-golang/pkg/user/handler"
)


func Setup(app *app.App) {
	app.HandleFunc("GET", "/login", handler.Login)
}
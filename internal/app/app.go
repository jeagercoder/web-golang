package app


import (
	"net/http"
	"fmt"
)



type App struct {
	routes map[string]map[string]http.HandlerFunc
	middlewares []http.HandlerFunc
}


func New() *App {
	return &App{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (app *App) Middleware(middleware http.HandlerFunc) {
	app.middlewares = append(app.middlewares, middleware)
}

func (app *App) HandleFunc(method, pattern string, handler http.HandlerFunc) {
	if _, exist := app.routes[method]; !exist {
		app.routes[method] = make(map[string]http.HandlerFunc)
	}
	app.routes[method][pattern] = handler
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path

	for _, middleware := range app.middlewares {
		middleware(w, r)
	}
	if handlers, ok := app.routes[method]; ok {
		if handler, ok := handlers[path]; ok {
			handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func (app *App) RunServer() {
	fmt.Println("Server is listening on port 8000")
	http.ListenAndServe(":8000", app)
}
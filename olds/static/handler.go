package static

import (
	"fmt"
	"net/http"
)

// define request handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// use a map to pass the handler
// key method-url
// value handler
type Engine struct {
	router map[string]HandlerFunc
}

// Create new Struct instance
// when add a use, dont forget to init, or it would panic
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// implement interface, make sure engine a handler (be used by ListenAndServe)
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// warp ListenAndServe()
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// make add request type easy
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// add post type request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// add get type request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

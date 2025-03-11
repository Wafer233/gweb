package gweb

import (
	"fmt"
	"net/http"
)

// define request handler
type HandlerFuncv1 func(http.ResponseWriter, *http.Request)

// use a map to pass the handler
type Enginev1 struct {
	router map[string]HandlerFuncv1
}

// Create new Struct instance
// when add a use, dont forget to init, or it would panic 3.11
func Newv1() *Enginev1 {
	return &Enginev1{router: make(map[string]HandlerFuncv1)}
}

// implement interface, make sure engine a handler (be used by ListenAndServe)
func (engine *Enginev1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// warp ListenAndServe()
func (engine *Enginev1) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// make add request type easy
func (engine *Enginev1) add(method string, pattern string, handler HandlerFuncv1) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// add post type request
func (engine *Enginev1) POST(pattern string, handler HandlerFuncv1) {
	engine.add("POST", pattern, handler)
}

// add get type request
func (engine *Enginev1) GET(pattern string, handler HandlerFuncv1) {
	engine.add("GET", pattern, handler)
}

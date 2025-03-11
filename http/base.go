package http

import (
	"fmt"
	"net/http"
)

// router
func withoutHandler() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

// handler echoes r.URL.Path
// handler func(ResponseWriter, *Request)
func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// Engine is the uni handler for all requests
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func WithHandler() {
	engine := new(Engine)
	http.ListenAndServe("localhost:8080", engine)
}

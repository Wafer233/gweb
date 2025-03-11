package middleware

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func onlyForGroup() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group ", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func Test_middleware(t *testing.T) {
	r := New()

	r.UseMiddleware(Logger()) // global midlleware
	r.GET("/hello", func(c *Context) {
		c.SendHTMLResponse(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/group")
	v2.UseMiddleware(onlyForGroup()) // v2 group middleware
	v2.GET("/hello/:name", func(c *Context) {
		// expect /hello/wafer
		c.SendTextResponse(http.StatusOK, "hello %s, you're at %s\n", c.GetQueryParam("name"), c.Path)
	})

	r.Run("localhost:8080")
}

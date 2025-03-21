package dynamic

import (
	"net/http"
	"testing"
)

func Test_handler(t *testing.T) {

	r := New()
	r.GET("/", func(c *Context) {
		c.SendHTMLResponse(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *Context) {
		// expect /hello?name=wafer
		c.SendTextResponse(http.StatusOK, "hello %s, you're at %s\n", c.GetQueryParam("name"), c.Path)
	})
	// expect /param/:words/param
	r.GET("/param/:words/param", func(c *Context) {
		c.SendTextResponse(http.StatusOK, "%s", c.GetDynamicParam("words"))
	})
	// expect /wildcard/gWeb.go
	r.GET("/wildcard/*filepath", func(c *Context) {
		c.SendTextResponse(http.StatusOK, "filepath: %s", c.GetDynamicParam("filepath"))
	})

	r.Run("localhost:8080")

}

package group

import (
	"net/http"
	"testing"
)

func Test_handler(t *testing.T) {

	r := New()

	r.GET("/", func(c *Context) {
		c.SendHTMLResponse(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	group := r.Group("/group")
	group.GET("/hello", func(c *Context) {
		// expect /group/hello?name=wafer
		c.SendTextResponse(http.StatusOK, "Hello, gWeb!")
	})
	group.GET("/hello/:name", func(c *Context) {
		// expect /group/hello/wafer
		c.SendTextResponse(http.StatusOK, "hello %s, you're at %s\n", c.GetDynamicParam("name"), c.Path)
	})
	group.GET("/assets/*filepath", func(c *Context) {
		c.SendJSONResponse(http.StatusOK, Hash{"filepath": c.GetDynamicParam("filepath")})
	})

	r.Run("localhost:8080")

}

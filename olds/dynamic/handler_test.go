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

	r.GET("/hello/:name", func(c *Context) {
		// expect /hello/wafer
		c.SendTextResponse(http.StatusOK, "hello %s, you're at %s\n", c.GetDynamicParam("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *Context) {
		c.SendJSONResponse(http.StatusOK, Hash{"filepath": c.GetDynamicParam("filepath")})
	})

	r.Run("localhost:8080")

}

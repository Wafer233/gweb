package context

import (
	"net/http"
	"testing"
)

// test hander v2 adding context/ router
func Test_handler(t *testing.T) {
	r := New()

	r.GET("/", func(c *Context) {
		c.SendHTMLResponse(http.StatusOK, "<h1> Hello Gweb </h1>")
	})

	r.GET("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.SendTextResponse(http.StatusOK, "hello %s, you're at the path: %s \n", c.GetQueryParam("name"), c.Path)
	})

	r.POST("/login", func(c *Context) {
		c.SendJSONResponse(http.StatusOK, Hash{
			"username": c.GetFormValue("username"),
			"password": c.GetFormValue("password"),
		})
	})

	r.Run("localhost:8080")
}

package context

import (
	"net/http"
	"testing"
)

// test hander v2 adding context/ router
func Test_handler(t *testing.T) {
	r := New()

	r.GET("/context", func(c *Context) {
		c.SendTextResponse(http.StatusOK, "Hello, gWeb!")
	})

	r.GET("/h", func(c *Context) {
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

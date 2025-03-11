package gweb

import (
	"net/http"
	"testing"
)

// test hander v2 adding context/ router
func Test_handler_v2(t *testing.T) {
	r := New()

	r.GET("/", func(c *Context) {
		c.SetHTML(http.StatusOK, "<h1> Hello Gweb </h1>")
	})

	r.GET("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.SetText(http.StatusOK, "hello %s, you're at the path: %s \n", c.GetQueryParam("name"), c.Path)
	})

	r.POST("/login", func(c *Context) {
		c.SetJSON(http.StatusOK, Hash{
			"username": c.GetBodyForm("username"),
			"password": c.GetBodyForm("password"),
		})
	})

	r.Run("localhost:8080")
}

package recovery

import (
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
	r := Default()
	r.GET("/", func(c *Context) {
		c.SendTextResponse(http.StatusOK, "Hello Wafer\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Context) {
		names := []string{"wafer"}
		c.SendTextResponse(http.StatusOK, names[100])
	})

	r.Run("localhost:8080")
}

func Test2(t *testing.T) {
	r := New()
	r.GET("/", func(c *Context) {
		c.SendTextResponse(http.StatusOK, "Hello Wafer\n")
	})

	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Context) {
		names := []string{"wafer"}
		c.SendTextResponse(http.StatusOK, names[100])
	})

	r.Run("localhost:8080")
}

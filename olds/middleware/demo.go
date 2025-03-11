package middleware

import "net/http"

// suppose middleware A B
// part1 -> part3 -> Handler -> part 4 -> part2
func A(c *Context) {
	c.SendTextResponse(http.StatusOK, "part1")
	c.Next()
	c.SendTextResponse(http.StatusOK, "part1")
}
func B(c *Context) {
	c.SendTextResponse(http.StatusOK, "part3")
	c.Next()
	c.SendTextResponse(http.StatusOK, "part4")
}

package main

import (
	"gweb"
	"net/http"
)

func main() {

	route := gweb.Default()
	route.POST("/", func(ctx *gweb.Context) {
		ctx.SendTextResponse(http.StatusOK, "hello world")
	})
	route.Run("localhost:8080")

}

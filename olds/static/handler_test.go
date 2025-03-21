package static

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_handler(t *testing.T) {
	r := New()
	r.GET("/static", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello, gWeb!")
	})

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run("localhost:8080")
}

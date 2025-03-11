package gweb

import (
	"fmt"
	"net/http"
	"testing"
)

// test hander
func Test_handler(t *testing.T) {
	r := Newv1()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run("localhost:8080")
}

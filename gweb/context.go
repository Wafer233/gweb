package gweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string

	// response info
	StatusCode int
}

// constructor
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// postman req: body
func (c *Context) GetBodyForm(key string) string {
	return c.Req.FormValue(key)
}

// postman req: URL-encoded query
func (c *Context) GetQueryParam(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// postman resp: URL-encoded query
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) SetText(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text")
	c.SetStatus(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) SetJSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//func (c *Context) SetData(code int, data []byte) {
//	c.SetStatus(code)
//	c.Writer.Write(data)
//}

func (c *Context) SetHTML(code int, html string) {
	c.SetHeader("Content-Type", "html")
	c.SetStatus(code)
	c.Writer.Write([]byte(html))
}

type Hash map[string]interface{}

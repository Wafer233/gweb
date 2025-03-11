package context

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
func (c *Context) GetFormValue(key string) string {
	return c.Req.FormValue(key)
}

// postman req: URL-encoded query
func (c *Context) GetQueryParam(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// postman resp: URL-encoded query
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) SendTextResponse(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text")
	c.SetStatusCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) SendJSONResponse(code int, obj interface{}) {
	c.SetHeader("Content-Type", "json")
	c.SetStatusCode(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) SendHTMLResponse(code int, html string) {
	c.SetHeader("Content-Type", "html")
	c.SetStatusCode(code)

	c.Writer.Write([]byte(html))
}

type Hash map[string]interface{}

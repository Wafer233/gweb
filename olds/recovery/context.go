package recovery

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// keyword
// Parameter matching :	e.g.	/p/:lang/doc	->	/p/c/doc 和 /p/go/doc。
// Wildcard * 			e.g		/static/*filepath	->	/static/fav.ico
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string

	// response info
	StatusCode int

	//add
	Params map[string]string

	// middleware
	handlers []HandlerFunc
	index    int
}

// constructor
// dont forget to init!!!!!!!!!!!!!!!!!!!
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,

		index: -1,
	}
}

// postman req: body
func (c *Context) GetFormValue(key string) string {
	return c.Req.FormValue(key)
}

// postman req: URL-encoded query
//update

func (c *Context) GetQueryParam(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) GetDynamicParam(key string) string {
	value, _ := c.Params[key]
	return value
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

// 当在中间件中调用Next方法时，控制权交给了下一个中间件，直到调用到最后一个中间件，
// 然后再从后往前，调用每个中间件在Next方法之后定义的部分。
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.SendJSONResponse(code, Hash{"message": err})
}

package minigin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// original
	Writer http.ResponseWriter
	Req    *http.Request
	// request
	Method string
	Path   string
	Paras  map[string]string
	//response
	StatusCode int

	//middlewares
	cntIndex int
	handlers []HandlerFunc
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:   w,
		Req:      req,
		Method:   req.Method,
		Paras:    make(map[string]string),
		Path:     req.URL.Path,
		cntIndex: -1,
	}
}

func (c *Context) CallNextHandler() {
	c.cntIndex++
	if c.cntIndex < len(c.handlers) {
		c.handlers[c.cntIndex](c)
	}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

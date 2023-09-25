package gee

import (
	"fmt"
	"net/http"
	"encoding/json"
)
type H map[string]interface{}

// Context is the encapsulation of http.Request and http.ResponseWriter
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

// NewContext is the constructor of Context
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm is the encapsulation of r.PostForm
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query is the encapsulation of r.URL.Query().Get()
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status sets the status code for response writer
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets the header for response writer
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

/* ------------------------------------------------------------- */

//Fast create response for String, JSON, Data, HTML
// String sets the string for response writer
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON sets the json for response writer
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil{
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data sets the data for response writer
func (c *Context) Data(code int, data []byte) {
	c.SetHeader("Content-Type", "application/octet-stream")
	c.Status(code)
	c.Writer.Write(data)
}

// HTML sets the html for response writer
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
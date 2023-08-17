package minigin

import (
	"net/http"
	"strings"
)

type Router struct {
	root *node
}

func newRouter() *Router {
	return &Router{root: newNode("GET", "/", false)}
}

func (r *Router) parsePattern(pattern string) []string {
	parts := strings.Split(pattern, "/")
	parts[0] = "/"
	return parts
}

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)
	updateTrie(r.root, method, parts, handler)
}

func (r *Router) handle(c *Context) {
	parts := r.parsePattern(c.Path)
	handler := searchTrie(0, r.root, c.Method, parts)
	if handler != nil {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

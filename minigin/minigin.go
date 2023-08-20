package minigin

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type (
	RouterGroup struct {
		prefix      string
		engine      *Engine
		middlewares []HandlerFunc
	}
	Engine struct {
		*RouterGroup
		router *Router
		groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.RegisterMiddleware(PanicRecovery(), Logger())
	return engine
}

func (e *Engine) NewRouterGroup(p string) *RouterGroup {
	newgroup := &RouterGroup{
		prefix: p,
		engine: e,
	}
	// fmt.Println(p)
	e.groups = append(e.groups, newgroup)
	return newgroup
}

func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) RegisterMiddleware(mdws ...HandlerFunc) {
	group.middlewares = append(group.middlewares, mdws...)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	middlewares := make([]HandlerFunc, 0)
	for _, grp := range e.groups {
		if strings.HasPrefix(req.URL.Path, grp.prefix) {
			middlewares = append(middlewares, grp.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	e.router.handle(c)
}

package minigin

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type (
	RouterGroup struct {
		Prefix      string
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

func (e *Engine) NewRouterGroup(p string) *RouterGroup {
	newgroup := &RouterGroup{
		Prefix: p,
		engine: e,
	}
	// fmt.Println(p)
	e.groups = append(e.groups, newgroup)
	return newgroup
}

func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = group.Prefix + pattern
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

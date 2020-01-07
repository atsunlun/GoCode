package gee

import (
	"net/http"
	"strings"
)

// 用来定义路由映射的处理方法
type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}

	// 路由
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// 创建路由
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 添加路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := getNestPrefix(group.parent, group.prefix) + comp
	group.engine.router.addRoute(method, pattern, handler)
}

// GET方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

func getNestPrefix(group *RouterGroup, p string) string {
	p = strings.Join([]string{group.prefix, p}, "")
	if group.parent == nil {
		return p
	}
	return getNestPrefix(group.parent, p)
}

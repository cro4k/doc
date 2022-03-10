package docer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type Engine struct {
	RouterGroup
}

func Wrap(e *gin.Engine) *Engine {
	v := reflect.ValueOf(e).Elem()
	engine := &Engine{}
	engine.e = e
	engine.basePath = v.FieldByName("basePath").String()
	engine.m = newManager()
	return engine
}

func (e *Engine) path(path string) string {
	return strings.TrimRight(e.basePath, "/") + "/" + strings.TrimLeft(path, "/")
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.e.(*gin.Engine).ServeHTTP(w, r)
}

func (e *Engine) Docs() *Manager {
	return e.m
}

type RouterGroup struct {
	e        gin.IRouter
	basePath string
	m        *Manager
}

func (e *RouterGroup) Use(h ...gin.HandlerFunc) *RouterGroup {
	e.e.Use(h...)
	return e
}

func (e *RouterGroup) Handle(method string, path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	handlers := []gin.HandlerFunc{handler.Handler}
	handlers = append(handlers, middle...)
	e.e.Handle(method, path, handlers...)
	handler.SetAPI(e.path(path))
	handler.SetMethod(method)
	e.m.add(handler)
	return e
}

func (e *RouterGroup) Any(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	handlers := []gin.HandlerFunc{handler.Handler}
	handlers = append(handlers, middle...)
	e.e.Any(path, handlers...)
	handler.SetAPI(e.path(path))
	e.m.add(handler)
	return e
}

func (e *RouterGroup) GET(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodGet, path, handler, middle...)
	return e
}
func (e *RouterGroup) POST(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodPost, path, handler, middle...)
	return e
}
func (e *RouterGroup) DELETE(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodDelete, path, handler, middle...)
	return e
}
func (e *RouterGroup) PATCH(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodPatch, path, handler, middle...)
	return e
}
func (e *RouterGroup) PUT(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodPut, path, handler, middle...)
	return e
}
func (e *RouterGroup) OPTIONS(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodOptions, path, handler, middle...)
	return e
}

func (e *RouterGroup) HEAD(path string, handler *Doc, middle ...gin.HandlerFunc) *RouterGroup {
	e.Handle(http.MethodHead, path, handler, middle...)
	return e
}

func (e *RouterGroup) Group(path string, handler ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{
		e:        e.e.Group(path, handler...),
		basePath: e.path(path),
		m:        e.m,
	}
}

func (e *RouterGroup) path(path string) string {
	return strings.TrimRight(e.basePath, "/") + "/" + strings.TrimLeft(path, "/")
}

func (e *RouterGroup) StaticFile(a string, b string) *RouterGroup {
	e.e.StaticFile(a, b)
	return e
}
func (e *RouterGroup) Static(a string, b string) *RouterGroup {
	e.e.Static(a, b)
	return e
}
func (e *RouterGroup) StaticFS(a string, b http.FileSystem) *RouterGroup {
	e.e.StaticFS(a, b)
	return e
}

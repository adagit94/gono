package goter

import (
	"github.com/adagit94/gono/gotils"
	"strings"
)

const (
	Post    = "POST"
	Get     = "GET"
	Put     = "PUT"
	Patch   = "PATCH"
	Delete  = "DELETE"
	Options = "OPTIONS"
	Connect = "CONNECT"
	Head    = "HEAD"
	Trace   = "TRACE"
)

type PathParams map[string]string

type IRouter interface {
	Route(path string) IRoute
	Handle(path string, method string)
}

func CreateRouter() IRouter {
	router := router{tree: make(tree)}
	return &router
}

type routeHandler func(pathParams PathParams)
type routes map[int][]routeConf
type tree map[string]routes

type segmentConf struct {
	segment string
	static  bool
}

type routeConf struct {
	segments []segmentConf
	handler  routeHandler
}

type router struct {
	tree tree
}

func (router *router) registerHandler(path string, method string, handler routeHandler) {
	segs := strings.Split(path, "/")
	segsCount := len(segs)

	if _, methodKeyExists := router.tree[method]; !methodKeyExists {
		router.tree[method] = make(routes)
	}

	if _, segsCountKeyExists := router.tree[method][segsCount]; !segsCountKeyExists {
		router.tree[method][segsCount] = make([]routeConf, 0)
	}

	router.tree[method][segsCount] = append(router.tree[method][segsCount], routeConf{segments: genSegConfs(segs), handler: handler})
	sortRoutes(router.tree[method][segsCount])
}

func (router *router) Route(path string) IRoute {
	route := route{path: path, registerHandler: router.registerHandler}
	return &route
}

func (router *router) Handle(path string, method string) {
	segs := strings.Split(path, "/")
	segsCount := len(segs)
	segsCountsMap, methodKey := router.tree[method]

	if !methodKey {
		panic(&gotils.CodeError{Code: gotils.MethodNotRegisteredCode, Message: "Method not registered."})
	}

	routes, segsCountKey := segsCountsMap[segsCount]

	if !segsCountKey {
		panic(&gotils.CodeError{Code: gotils.RouteNotRegisteredCode, Message: "Route not registered."})
	}

	for _, routeConf := range routes {
		pathParams := make(PathParams)
		take := true

		for i, seg := range routeConf.segments {
			if seg.static {
				if seg.segment != segs[i] {
					take = false
					break
				}
			} else {
				pathParams[seg.segment] = segs[i]
			}
		}

		if take {
			routeConf.handler(pathParams)
			break
		}
	}
}

type route struct {
	path            string
	registerHandler func(path string, method string, handler routeHandler)
}

type IRoute interface {
	Post(routeHandler) IRoute
	Get(routeHandler) IRoute
	Put(routeHandler) IRoute
	Patch(routeHandler) IRoute
	Delete(routeHandler) IRoute
	Options(routeHandler) IRoute
	Connect(routeHandler) IRoute
	Head(routeHandler) IRoute
	Trace(routeHandler) IRoute
}

func (route *route) Post(handler routeHandler) IRoute {
	route.registerHandler(route.path, Post, handler)
	return route
}

func (route *route) Get(handler routeHandler) IRoute {
	route.registerHandler(route.path, Get, handler)
	return route
}

func (route *route) Put(handler routeHandler) IRoute {
	route.registerHandler(route.path, Put, handler)
	return route
}

func (route *route) Patch(handler routeHandler) IRoute {
	route.registerHandler(route.path, Patch, handler)
	return route
}

func (route *route) Delete(handler routeHandler) IRoute {
	route.registerHandler(route.path, Delete, handler)
	return route
}

func (route *route) Options(handler routeHandler) IRoute {
	route.registerHandler(route.path, Options, handler)
	return route
}

func (route *route) Connect(handler routeHandler) IRoute {
	route.registerHandler(route.path, Connect, handler)
	return route
}

func (route *route) Head(handler routeHandler) IRoute {
	route.registerHandler(route.path, Head, handler)
	return route
}

func (route *route) Trace(handler routeHandler) IRoute {
	route.registerHandler(route.path, Trace, handler)
	return route
}

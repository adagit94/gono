package goter

import (
	"github.com/adagit94/gono/gotils"
)

type statRouteHandler func()
type statRoutes map[string]statRouteHandler
type statTree map[string]statRoutes

type IStatRouter interface {
	Route(path string) IStatRoute
	Handle(path string, method string)
}

func CreateStaticRouter() IStatRouter {
	router := statRouter{tree: make(statTree)}
	return &router
}

type statRouter struct {
	tree statTree
}

func (router *statRouter) registerHandler(path string, method string, handler statRouteHandler) {
	_, keyExists := router.tree[method]

	if !keyExists {
		router.tree[method] = make(statRoutes)
	}

	router.tree[method][path] = handler
}

func (router *statRouter) Route(path string) IStatRoute {
	r := statRoute{path: path, registerHandler: router.registerHandler}
	return &r
}

func (router *statRouter) Handle(path string, method string) {
	routes, methodKey := router.tree[method]

	if !methodKey {
		panic(&gotils.CodeError{Code: gotils.MethodNotRegisteredCode, Message: "Method not registered."})
	}

	handler, routeKey := routes[path]

	if !routeKey {
		panic(&gotils.CodeError{Code: gotils.RouteNotRegisteredCode, Message: "Route not registered."})
	}

	handler()
}

type statRoute struct {
	path            string
	registerHandler func(path string, method string, handler statRouteHandler)
}

type IStatRoute interface {
	Post(statRouteHandler) IStatRoute
	Get(statRouteHandler) IStatRoute
	Put(statRouteHandler) IStatRoute
	Patch(statRouteHandler) IStatRoute
	Delete(statRouteHandler) IStatRoute
	Options(statRouteHandler) IStatRoute
	Connect(statRouteHandler) IStatRoute
	Head(statRouteHandler) IStatRoute
	Trace(statRouteHandler) IStatRoute
}

func (route *statRoute) Post(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Post, handler)
	return route
}

func (route *statRoute) Get(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Get, handler)
	return route
}

func (route *statRoute) Put(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Put, handler)
	return route
}

func (route *statRoute) Patch(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Patch, handler)
	return route
}

func (route *statRoute) Delete(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Delete, handler)
	return route
}

func (route *statRoute) Options(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Options, handler)
	return route
}

func (route *statRoute) Connect(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Connect, handler)
	return route
}

func (route *statRoute) Head(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Head, handler)
	return route
}

func (route *statRoute) Trace(handler statRouteHandler) IStatRoute {
	route.registerHandler(route.path, Trace, handler)
	return route
}

package router

import (
	"github.com/valyala/fasthttp"
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

type ReqCtx = *fasthttp.RequestCtx
type PathParams map[string]string

type IRoute interface {
	Post(reqHandler) IRoute
	Get(reqHandler) IRoute
	Put(reqHandler) IRoute
	Patch(reqHandler) IRoute
	Delete(reqHandler) IRoute
	Options(reqHandler) IRoute
	Connect(reqHandler) IRoute
	Head(reqHandler) IRoute
	Trace(reqHandler) IRoute
}

type IRouter interface {
	Route(path string) IRoute
	HandleReq(ctx ReqCtx)
}

func CreateRouter() IRouter {
	router := router{tree: make(tree)}

	return &router
}

type reqHandler func(ctx ReqCtx, pathParams PathParams)
type routes map[int][]routeConf
type tree map[string]routes

type segmentConf struct {
	segment string
	static  bool
}

type routeConf struct {
	segments []segmentConf
	handler  reqHandler
}

type router struct {
	tree tree
}

func (router *router) registerHandler(path string, method string, handler reqHandler) {
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
	r := route{path: path, registerHandler: router.registerHandler}
	return &r
}

func (router *router) HandleReq(ctx ReqCtx) {
	segs := strings.Split(unsafeString(ctx.Path()), "/")
	segsCount := len(segs)
	segsCountsMap, methodKey := router.tree[unsafeString(ctx.Method())]

	if !methodKey {
		methodNotFound(ctx)
		return
	}

	routes, segsCountKey := segsCountsMap[segsCount]

	if !segsCountKey {
		routeNotFound(ctx)
		return
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
			routeConf.handler(ctx, pathParams)
			break
		}
	}
}

type route struct {
	path            string
	registerHandler func(path string, method string, handler reqHandler)
}

func (route *route) Post(handler reqHandler) IRoute {
	route.registerHandler(route.path, Post, handler)
	return route
}

func (route *route) Get(handler reqHandler) IRoute {
	route.registerHandler(route.path, Get, handler)
	return route
}

func (route *route) Put(handler reqHandler) IRoute {
	route.registerHandler(route.path, Put, handler)
	return route
}

func (route *route) Patch(handler reqHandler) IRoute {
	route.registerHandler(route.path, Patch, handler)
	return route
}

func (route *route) Delete(handler reqHandler) IRoute {
	route.registerHandler(route.path, Delete, handler)
	return route
}

func (route *route) Options(handler reqHandler) IRoute {
	route.registerHandler(route.path, Options, handler)
	return route
}

func (route *route) Connect(handler reqHandler) IRoute {
	route.registerHandler(route.path, Connect, handler)
	return route
}

func (route *route) Head(handler reqHandler) IRoute {
	route.registerHandler(route.path, Head, handler)
	return route
}

func (route *route) Trace(handler reqHandler) IRoute {
	route.registerHandler(route.path, Trace, handler)
	return route
}

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

type statReqHandler func(ctx *fasthttp.RequestCtx)
type StatMethods map[string]statReqHandler
type statRoutes map[string]statReqHandler
type statTree map[string]statRoutes

type statRouter struct {
	tree statTree
}

func (router *statRouter) AddRoute(path string, methods StatMethods) {
	for method, handler := range methods {
		_, keyExists := router.tree[method]

		if !keyExists {
			router.tree[method] = make(statRoutes)
		}

		router.tree[method][path] = handler
	}
}

func (router *statRouter) HandleReq(ctx ReqCtx) {
	routes, methodKey := router.tree[unsafeString(ctx.Method())]

	if !methodKey {
		methodNotFound(ctx)
		return
	}

	handler, routeKey := routes[unsafeString(ctx.Path())]

	if !routeKey {
		routeNotFound(ctx)
		return
	}

	handler(ctx)
}

type IStatRouter interface {
	AddRoute(path string, handlers StatMethods)
	HandleReq(ctx ReqCtx)
}

func CreateStaticRouter() IStatRouter {
	router := statRouter{tree: make(statTree)}

	return &router
}

type dynReqHandler func(ctx *fasthttp.RequestCtx, pathParams PathParams)
type DynMethods map[string]dynReqHandler
type dynRoutes map[int][]dynRouteConf
type dynTree map[string]dynRoutes

type segmentConf struct {
	segment string
	static  bool
}

type dynRouteConf struct {
	segments []segmentConf
	handler  dynReqHandler
}

type dynRouter struct {
	tree dynTree
}

func (router *dynRouter) AddRoute(path string, methods DynMethods) {
	segs := strings.Split(path, "/")
	segsCount := len(segs)

	for method, handler := range methods {
		if _, methodKeyExists := router.tree[method]; !methodKeyExists {
			router.tree[method] = make(dynRoutes)
		}

		if _, segsCountKeyExists := router.tree[method][segsCount]; !segsCountKeyExists {
			router.tree[method][segsCount] = make([]dynRouteConf, 0)
		}

		router.tree[method][segsCount] = append(router.tree[method][segsCount], dynRouteConf{segments: genSegConfs(segs), handler: handler})
		sortRoutes(router.tree[method][segsCount])
	}

}

func (router *dynRouter) HandleReq(ctx ReqCtx) {
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

type IDynRouter interface {
	AddRoute(path string, handlers DynMethods)
	HandleReq(ctx ReqCtx)
}

func CreateRouter() IDynRouter {
	router := dynRouter{tree: make(dynTree)}

	return &router
}

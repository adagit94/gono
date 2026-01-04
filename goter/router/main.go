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
type RegisterHandler[T StatReqHandler | DynReqHandler] func(handler T) *MethodsHandlers[T]
type MethodsHandlers[T StatReqHandler | DynReqHandler] struct {
	Post, Get, Put, Patch, Delete, Options, Connect, Head, Trace RegisterHandler[T]
}

type StatReqHandler func(ctx ReqCtx)
type StatMethods map[string]StatReqHandler
type statRoutes map[string]StatReqHandler
type statTree map[string]statRoutes

type statRouter struct {
	tree statTree
}

func (router *statRouter) Route(path string) *MethodsHandlers[StatReqHandler] {
	routeMethods := MethodsHandlers[StatReqHandler]{}
	addMethodHandlerRegistrator := func(method string) func(handler StatReqHandler) *MethodsHandlers[StatReqHandler] {
		return func(handler StatReqHandler) *MethodsHandlers[StatReqHandler] {
			_, keyExists := router.tree[method]

			if !keyExists {
				router.tree[method] = make(statRoutes)
			}

			router.tree[method][path] = handler

			return &routeMethods
		}
	}

	routeMethods.Post = addMethodHandlerRegistrator(Post)
	routeMethods.Get = addMethodHandlerRegistrator(Get)
	routeMethods.Put = addMethodHandlerRegistrator(Put)
	routeMethods.Patch = addMethodHandlerRegistrator(Patch)
	routeMethods.Delete = addMethodHandlerRegistrator(Delete)
	routeMethods.Options = addMethodHandlerRegistrator(Options)
	routeMethods.Connect = addMethodHandlerRegistrator(Connect)
	routeMethods.Head = addMethodHandlerRegistrator(Head)
	routeMethods.Trace = addMethodHandlerRegistrator(Trace)

	return &routeMethods

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
	Route(path string) *MethodsHandlers[StatReqHandler]
	HandleReq(ctx ReqCtx)
}

func CreateStaticRouter() IStatRouter {
	router := statRouter{tree: make(statTree)}

	return &router
}

type DynReqHandler func(ctx ReqCtx, pathParams PathParams)
type DynMethods map[string]DynReqHandler
type dynRoutes map[int][]dynRouteConf
type dynTree map[string]dynRoutes

type segmentConf struct {
	segment string
	static  bool
}

type dynRouteConf struct {
	segments []segmentConf
	handler  DynReqHandler
}

type dynRouter struct {
	tree dynTree
}

func (router *dynRouter) Route(path string) *MethodsHandlers[DynReqHandler] {
	routeMethods := MethodsHandlers[DynReqHandler]{}
	addMethodHandlerRegistrator := func(method string) func(handler DynReqHandler) *MethodsHandlers[DynReqHandler] {
		return func(handler DynReqHandler) *MethodsHandlers[DynReqHandler] {
			segs := strings.Split(path, "/")
			segsCount := len(segs)

			if _, methodKeyExists := router.tree[method]; !methodKeyExists {
				router.tree[method] = make(dynRoutes)
			}

			if _, segsCountKeyExists := router.tree[method][segsCount]; !segsCountKeyExists {
				router.tree[method][segsCount] = make([]dynRouteConf, 0)
			}

			router.tree[method][segsCount] = append(router.tree[method][segsCount], dynRouteConf{segments: genSegConfs(segs), handler: handler})
			sortRoutes(router.tree[method][segsCount])

			return &routeMethods
		}
	}

	routeMethods.Post = addMethodHandlerRegistrator(Post)
	routeMethods.Get = addMethodHandlerRegistrator(Get)
	routeMethods.Put = addMethodHandlerRegistrator(Put)
	routeMethods.Patch = addMethodHandlerRegistrator(Patch)
	routeMethods.Delete = addMethodHandlerRegistrator(Delete)
	routeMethods.Options = addMethodHandlerRegistrator(Options)
	routeMethods.Connect = addMethodHandlerRegistrator(Connect)
	routeMethods.Head = addMethodHandlerRegistrator(Head)
	routeMethods.Trace = addMethodHandlerRegistrator(Trace)

	return &routeMethods

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
	Route(path string) *MethodsHandlers[DynReqHandler]
	HandleReq(ctx ReqCtx)
}

func CreateRouter() IDynRouter {
	router := dynRouter{tree: make(dynTree)}

	return &router
}

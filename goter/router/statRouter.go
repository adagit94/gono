package router

type statReqHandler func(ctx ReqCtx)
type statRoutes map[string]statReqHandler
type statTree map[string]statRoutes

type IStatRoute interface {
	Post(statReqHandler) IStatRoute
	Get(statReqHandler) IStatRoute
	Put(statReqHandler) IStatRoute
	Patch(statReqHandler) IStatRoute
	Delete(statReqHandler) IStatRoute
	Options(statReqHandler) IStatRoute
	Connect(statReqHandler) IStatRoute
	Head(statReqHandler) IStatRoute
	Trace(statReqHandler) IStatRoute
}

type IStatRouter interface {
	Route(path string) IStatRoute
	HandleReq(ctx ReqCtx)
}

func CreateStaticRouter() IStatRouter {
	router := statRouter{tree: make(statTree)}

	return &router
}

type statRouter struct {
	tree statTree
}

func (router *statRouter) registerHandler(path string, method string, handler statReqHandler) {
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

type statRoute struct {
	path            string
	registerHandler func(path string, method string, handler statReqHandler)
}

func (route *statRoute) Post(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Post, handler)
	return route
}

func (route *statRoute) Get(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Get, handler)
	return route
}

func (route *statRoute) Put(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Put, handler)
	return route
}

func (route *statRoute) Patch(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Patch, handler)
	return route
}

func (route *statRoute) Delete(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Delete, handler)
	return route
}

func (route *statRoute) Options(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Options, handler)
	return route
}

func (route *statRoute) Connect(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Connect, handler)
	return route
}

func (route *statRoute) Head(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Head, handler)
	return route
}

func (route *statRoute) Trace(handler statReqHandler) IStatRoute {
	route.registerHandler(route.path, Trace, handler)
	return route
}

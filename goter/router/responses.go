package router

func SendText(text string, ctx ReqCtx) {
	ctx.SetContentType("text/plain; charset=utf8")
	ctx.Response.AppendBodyString(text)
}

func NotFound(text string, ctx ReqCtx) {
	ctx.Response.SetStatusCode(404)
	SendText(text, ctx)
}

func methodNotFound(ctx ReqCtx) {
	NotFound("Method not registered.", ctx)
}

func routeNotFound(ctx ReqCtx) {
	NotFound("Route not registered.", ctx)
}

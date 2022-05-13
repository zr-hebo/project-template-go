package web

import (
	"net/http/pprof"

	"github.com/kataras/iris/v12"
)

func registerPprofHandler() {
	webService.Get("/debug/pprof/", outletPprofIndex)
	webService.Get("/debug/heap", outletPprofHeap)
	webService.Get("/debug/pprof/heap", outletPprofHeap)
	webService.Get("/debug/goroutine", outletPprofGoroutine)
	webService.Get("/debug/pprof/goroutine", outletPprofGoroutine)
	webService.Get("/debug/allocs", outletPprofAllocs)
	webService.Get("/debug/pprof/allocs", outletPprofAllocs)
	webService.Get("/debug/block", outletPprofBlock)
	webService.Get("/debug/pprof/block", outletPprofBlock)
	webService.Get("/debug/threadcreate", outletPprofThreadCreate)
	webService.Get("/debug/pprof/threadcreate", outletPprofThreadCreate)
	webService.Get("/debug/cmdline", outletPprofCmdline)
	webService.Get("/debug/pprof/cmdline", outletPprofCmdline)
	webService.Get("/debug/profile", outletPprofProfile)
	webService.Get("/debug/pprof/profile", outletPprofProfile)
	webService.Get("/debug/symbol", outletPprofSymbol)
	webService.Get("/debug/pprof/symbol", outletPprofSymbol)
	webService.Post("/debug/symbol", outletPprofSymbol)
	webService.Post("/debug/pprof/symbol", outletPprofSymbol)
	webService.Get("/debug/trace", outletPprofTrace)
	webService.Get("/debug/pprof/trace", outletPprofTrace)
	webService.Get("/debug/mutex", outletPprofMutex)
	webService.Get("/debug/pprof/mutex", outletPprofMutex)
}

func outletPprofIndex(ctx iris.Context) {
	pprof.Index(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofHeap(ctx iris.Context) {
	pprof.Handler("heap").ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofGoroutine(ctx iris.Context) {
	pprof.Handler("goroutine").ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofAllocs(ctx iris.Context) {
	pprof.Handler("allocs").ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofBlock(ctx iris.Context) {
	pprof.Handler("block").ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofThreadCreate(ctx iris.Context) {
	pprof.Handler("threadcreate").ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofCmdline(ctx iris.Context) {
	pprof.Cmdline(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofProfile(ctx iris.Context) {
	pprof.Profile(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofSymbol(ctx iris.Context) {
	pprof.Symbol(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofTrace(ctx iris.Context) {
	pprof.Trace(ctx.ResponseWriter(), ctx.Request())
}

func outletPprofMutex(ctx iris.Context) {
	pprof.Handler("mutex").ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

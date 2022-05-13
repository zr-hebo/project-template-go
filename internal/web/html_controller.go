package web

import (
	"os"

	log "github.com/golang/glog"
	"github.com/kataras/iris/v12"
)

const templatePath = "web/views"

func registerPage() {
	_, err := os.Stat(templatePath)
	if err != nil {
		if !os.IsExist(err) {
			log.Warningf("web template path:%s not exist", templatePath)
			webService.Get("/", func(ctx iris.Context) {
				_, _ = ctx.ResponseWriter().WriteString(indexPageContent)
			})
			return
		}
	}

	webService.RegisterView(iris.HTML(templatePath, ".html"))
	webService.Get("/", controllerRoot)
}

func controllerRoot(ctx iris.Context) {
	ctx.View("index.html")
}

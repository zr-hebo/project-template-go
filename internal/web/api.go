package web

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/kataras/iris/v12"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zr-hebo/project-template-go/config"
	hu "github.com/zr-hebo/utils/http"
)

var (
	statusRouter = mux.NewRouter()
	webService   = iris.New()
	promHandler  = promhttp.Handler()
)

func initEnv() {
	registerPage()
	registerStatusHandler()
	registerConfigHandler()
	registerPprofHandler()

	webService.WrapRouter(func(w http.ResponseWriter, r *http.Request, handler http.HandlerFunc) {
		path := r.URL.Path
		if path == "/debug/metrics" {
			promHandler.ServeHTTP(w, r)
		} else if strings.HasPrefix(path, "/status/") {
			statusRouter.ServeHTTP(w, r)
		} else {
			handler(w, r)
		}
	})
}

func StartWebService() {
	initEnv()
	listenAddr := fmt.Sprintf("0.0.0.0:%d", config.WebServicePort.MustGetVal())

	go func() {
		err := webService.Run(iris.Addr(listenAddr))
		if err != nil {
			log.Fatalf("open web service failed <-- %s", err.Error())
		}
	}()

	time.Sleep(time.Second)
	log.Infof("listen web service on %s OK", listenAddr)
}

func registerStatusHandler() {
}

func registerConfigHandler() {
}

func outletGetStatus(resp http.ResponseWriter, _ *http.Request) {
	mp := hu.NewMouthpiece(resp)
	mp.Data = "running"
	err := mp.Convey()
	if err != nil {
		log.Errorf("deal GetStatus request failed <-- %s", err.Error())
	}
}

package base

import (
	"adminbg/pkg/g"
	"adminbg/pkg/log"
	"adminbg/pkg/util"
	"adminbg/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type AdminBgServer struct {
	httpSrv *http.Server
}

func (a *AdminBgServer) Init() {
	log.Infoln("<------  ADMIN BG Initiating  ----->")
	ginEngine := gin.Default()
	router.Init(ginEngine)

	a.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", g.Conf.AppAdminbg.Host, g.Conf.AppAdminbg.Port),
		Handler: ginEngine,
	}
}

func (a *AdminBgServer) Run() {
	log.Infoln("<------  ADMIN BG is Running  ----->")
	httpSrvQuit := make(chan struct{})
	go func() {
		err := a.httpSrv.ListenAndServe()
		util.PanicIfErr(err, []error{http.ErrServerClosed})
		httpSrvQuit <- struct{}{}
	}()
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	select {
	case <-exit:
	case <-httpSrvQuit:
	}
}

func (a *AdminBgServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_ = a.httpSrv.Shutdown(ctx)
	cancel()

	g.Stop() // 注意：最后回收资源，否则会影响使用这些资源的goroutine
	log.Infoln("<------  ADMIN BG Exited  ----->")
}

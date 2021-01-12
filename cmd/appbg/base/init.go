package base

import (
	"adminbg/log"
	"adminbg/pkg/g"
	"adminbg/router"
	"adminbg/util"
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
	log.Infoln("<------------  ADMIN BG is initiating  ----------->")
	ginEngine := gin.Default()
	router.Init(ginEngine)

	a.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", g.Conf.AppAdminbg.Host, g.Conf.AppAdminbg.Port),
		Handler: ginEngine,
	}
}

func (a *AdminBgServer) Run() {
	log.Infoln("<------------  ADMIN BG is running  ----------->")
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

	g.Stop() // release global objects at last
	log.Infoln("<------------  ADMIN BG exited  ----------->")
}

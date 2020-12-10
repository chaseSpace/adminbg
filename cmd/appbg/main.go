package main

import (
	"adminbg/config"
	"adminbg/pkg/_util/_config"
	"adminbg/pkg/g"
	"adminbg/pkg/log"
	"adminbg/router"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	cfgFile = flag.String("cfgFile", "../../deploy/config/settings.dev.yml", "")
)

var (
	httpSrv *http.Server
)

func init() {
	flag.Parse()

	g.Conf = new(config.Conf)
	_config.MustLoadByFile(*cfgFile, g.Conf)
	//fmt.Printf("%+v", g.Conf)

	g.MustInit() // 所有全局资源初始化(db等)
	log.MustInit(g.Conf.Logger)
}

func gracefulStop() {
	var ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = httpSrv.Shutdown(ctx)
	/*
	 */
	g.Stop() // 注意：最后回收资源，否则会影响使用这些资源的线程
}

func main() {
	log.Println("-------- running")

	engine := gin.Default()
	router.Init(engine)

	httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", g.Conf.App.Host, g.Conf.App.Port),
		Handler: engine,
	}
	httpSrvExit := make(chan error)
	go func() {
		err := httpSrv.ListenAndServe()
		log.Println("gin router exited, err:", err)
		httpSrvExit <- err
	}()

	defer gracefulStop()

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	select {
	case <-exit:
	case <-httpSrvExit:
	}

	log.Println("-------- exited")
}

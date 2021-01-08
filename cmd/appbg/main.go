package main

import (
	"adminbg/cmd/appbg/base"
	"adminbg/config"
	"adminbg/pkg/g"
	"adminbg/pkg/log"
	"adminbg/pkg/util/_config"
	"flag"
	"fmt"
)

var (
	cfgFile = flag.String("cfgFile", "deploy/config/conf.dev.yml", "")
)

func init() {
	flag.Parse()

	g.Conf = new(config.Conf)
	_config.MustLoadByFile(*cfgFile, g.Conf)
	fmt.Printf("%+v\n", g.Conf)

	// 所有全局资源初始化(db等)
	g.MustInit()
	log.MustInit(g.Conf.Logger)
}

func main() {
	server := new(base.AdminBgServer)
	server.Init()
	defer server.Stop()
	server.Run()
}

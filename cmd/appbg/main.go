package main

import (
	"adminbg/cmd/appbg/base"
	"adminbg/config"
	"adminbg/log"
	"adminbg/pkg/g"
	"adminbg/pkg/model"
	"adminbg/util/_config"
	"flag"
)

var (
	cfgFile = flag.String("cfgFile", "deploy/config/conf.dev.yml", "")
)

func init() {
	flag.Parse()

	g.Conf = new(config.Conf)
	_config.MustLoadByFile(*cfgFile, g.Conf)
	//fmt.Printf("%+v\n", g.Conf)
	g.Conf.AssertOK()

	g.MustInit()
	model.MustInit(g.Mysql)

	log.MustInit(g.Conf.Logger)
}

func main() {
	server := new(base.AdminBgServer)
	server.Init()
	defer server.Stop()
	server.Run()
}

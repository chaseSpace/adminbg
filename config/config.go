package config

import (
	"adminbg/util"
	"time"
)

type Conf struct {
	AppAdminbg AppAdminbg `yaml:"app_adminbg"`
	Logger     Logger
	Mysql      Mysql
}

type AppAdminbg struct {
	Mode string // dev | test | prod
	Name string
	Host string
	Port int16
	Jwt  Jwt
}

type Logger struct {
	Dir               string
	Level             string
	DBLogFilename     string `yaml:"db_log_filename"`
	CommonLogFilename string `yaml:"common_log_filename"`
	ToStdout          bool   `yaml:"to_stdout"`
}

type Jwt struct {
	Secret        string
	Timeout       int32
	TimeoutForDev int32 `yaml:"timeout_for_dev"`
}

type Mysql struct {
	Source string
}

func (c *Conf) AssertOK() {
	if c.AppAdminbg.Jwt.Secret == "" {
		panic("Conf.AppAdminbg.Jwt.Secret is empty!")
	}
	timeout := time.Duration(c.AppAdminbg.Jwt.Timeout) * time.Second
	timeoutDev := time.Duration(c.AppAdminbg.Jwt.TimeoutForDev) * time.Second
	util.InitJWT(timeout, timeoutDev, 0, c.AppAdminbg.Jwt.Secret)
}

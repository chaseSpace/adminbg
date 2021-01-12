package config

/*
存放全局配置，注意必要时使用yaml tag，否则读取不到配置
*/

type Conf struct {
	AppAdminbg AppAdminbg `yaml:"app_adminbg"`
	Logger     Logger
	Mysql      Mysql
}

type AppAdminbg struct {
	Mode string
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
	Secret    string
	Timeout   int32
	TestToken string `yaml:"test_token"`
}

type Mysql struct {
	Source string
}

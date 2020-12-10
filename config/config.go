package config

/*
存放全局配置，注意必要时使用yaml tag，否则读取不到配置
*/

type Conf struct {
	App      App
	Logger   Logger
	Jwt      Jwt
	Database Database
	Others   Others
	Kafka    Kafka
}

type App struct {
	Mode         string
	Host         string
	Name         string
	Port         int16
	ReadTimeout  int8 `yaml:"read_timeout"`
	WriteTimeout int8 `yaml:"write_timeout"`
}

type Logger struct {
	Dir                string
	DBLogFilename      string `yaml:"db_log_filename"`
	RequestLogFilename string `yaml:"req_log_filename"`
	DefaultLogFilename string `yaml:"default_log_filename"`
	// ...
	ToStdout bool `yaml:"to_stdout"`
}

type Jwt struct {
	Secret    string
	Timeout   int32
	TestToken string `yaml:"test_token"`
}

type Database struct {
	Mysql      Mysql
	Clickhouse Clickhouse
}

type Mysql struct {
	Source string
}

type Clickhouse struct {
	Source string
}

type Others struct {
	PplApiAddr string `yaml:"ppl_api_addr"`
}

type Kafka struct {
	Hosts      string
	Topics     string
	ConsumerId string `yaml:"consumer_id"`
	GroupId    string `yaml:"group_id"`
	ClientId   string `yaml:"client_id"`
}

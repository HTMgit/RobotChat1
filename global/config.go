package global

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	Config   Options
	Location *time.Location
)

type Options struct {
	Server  server
	Logger  logger
	BaseCfg baseConfig
	Mysql   mysql
	ReqUrl  requrl
}

type baseConfig struct {
	Env       string
	EtcdAddrs []string
	TimeZone  string
}

type server struct {
	Port    string
	Host    string
	Mode    string
	Version string
}

type logger struct {
	OpLogPath         string
	ErrLogPath        string
	LogLevel          string
	LogFileMaxSize    int
	LogFileMaxBackups int
	LogFileMaxAge     int
}

type mysql struct {
	Address  string
	User     string
	Password string
	Database string
}

type requrl struct {
	Xiaoyingurl     string
	Liaotiannvpuurl string
	Englishurl      string
}

func LoadConfig(pathToToml string) {
	if _, err := toml.DecodeFile(pathToToml, &Config); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ReloadConfig() {
	filePath, err := filepath.Abs("./config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("parse toml file once. filePath: %s\n", filePath)
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("new Config = %+v \n", Config)

}

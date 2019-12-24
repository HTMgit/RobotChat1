package logger

import (
	"fmt"
	"os"
	"robot_chat/global"

	"github.com/lixuanhao/fileLogger"
)

var (
	Logger   *fileLogger.FileLogger
	CountLog *fileLogger.FileLogger
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func SetupLogger(appName string) {
	path := "./logs"
	if !isExist(path) {
		os.Mkdir(path, 0755)
	}
	fmt.Println("path:", path)
	fmt.Println("appName:", appName)
	Logger = fileLogger.NewDailyLogger(path, appName+".log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)

	Logger.SetLogConsole(false)
	switch global.Config.Logger.LogLevel {
	case "trace":
		Logger.SetLogLevel(fileLogger.TRACE)
	case "info":
		Logger.SetLogLevel(fileLogger.INFO)
	case "warn":
		Logger.SetLogLevel(fileLogger.WARN)
	case "error":
		Logger.SetLogLevel(fileLogger.ERROR)
	case "off":
		Logger.SetLogLevel(fileLogger.OFF)
	default:
		Logger.SetLogLevel(fileLogger.TRACE)
	}

}

/*初始化count模块logger*/
func SetupCountLogger(appName string) {
	var path string
	path = "./logs"

	if !isExist(path) {
		os.Mkdir(path, 0755)
	}
	CountLog = fileLogger.NewDailyLogger(path, appName+".data.log", "", fileLogger.DEFAULT_LOG_SCAN, fileLogger.DEFAULT_LOG_SEQ)
}

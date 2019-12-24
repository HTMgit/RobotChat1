package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"robot_chat/global"
	"robot_chat/logger"
	"robot_chat/routers"
	"robot_chat/store"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	wg         sync.WaitGroup
	configFile = flag.String("config", "./config.toml", "Path to toml config file.")
)

func main() {
	flag.Parse()

	fmt.Println("*configFile", *configFile)
	global.LoadConfig(*configFile)
	var err error
	global.Location, err = time.LoadLocation(global.Config.BaseCfg.TimeZone)
	if err != nil {
		panic(fmt.Sprintf("[Main] load location failed, err=%v", err))
	}

	// logger.ConfigureLogger(global.Config.Logger.ErrLogPath, global.Config.Logger.OpLogPath, fmt.Sprintf("%s:%s", os.Args[0], global.Config.Server.Port), global.Config.Server.Version)
	logger.SetupLogger("robot_chat")
	logger.SetupCountLogger("robot_chat")

	defer func() {
		err := recover()
		if err != nil {
			logger.Logger.Error("main panic error : %v", err)
		} else {
			logger.Logger.Info("main exit")
		}
	}()

	store.NewMysql(global.Config.Mysql.User, global.Config.Mysql.Password, global.Config.Mysql.Address, global.Config.Mysql.Database)
	defer store.CloseMysql()

	r := gin.New()
	routers.NewRouter(r)
	srv := &http.Server{
		Addr:    net.JoinHostPort(global.Config.Server.Host, global.Config.Server.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("\nlisten: %s\n", err.Error())
		}
	}()

	logger.Logger.Info("[Main]################################################# SERVER START ###################################################")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server Shutdown:%s\n", err.Error())
	}
	select {
	case <-ctx.Done():
		logger.Logger.Info("timeout of 5 seconds.")
	}
	wg.Wait()
	logger.Logger.Info("Server exit.")

}

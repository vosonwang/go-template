// 生成swagger配置
//go:generate swag init -g router/router.go
package main

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"log"
	"my-project-name/aliyun/acm"
	"my-project-name/aliyun/sls"
	"my-project-name/config"
	"my-project-name/db/mongodb"
	"my-project-name/logger"
	"my-project-name/router"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const project = "my-project-name"

func main() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 阿里云ACM应用配置管理
	acm.Init()

	sls.SetLogStore(project)

	config.Init(project, mongodb.Init, sls.Init)

	logger.Init(project, SLSHook)

	e := router.Init()

	logger.Info(fmt.Sprintf("service online. service:%v,host:%v,version:%v,env:%v,debug:%v",
		project, config.Host(), config.Version(), config.Env(), config.Debug()))

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func SLSHook(e zapcore.Entry) error {
	return sls.SendLog(e.Level.String(), config.Host(), e.Time, map[string]string{
		"msg":  e.Message,
		"file": e.Caller.TrimmedPath(),
	})
}

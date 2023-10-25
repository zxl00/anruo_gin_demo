package cmd

import (
	"anruo-gin-demo/common"
	"anruo-gin-demo/routers"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("server.address"),
			viper.GetInt("server.port")),
		Handler:        routers.BaseRouters(),
		MaxHeaderBytes: 1 << 20,
	}
	// 打印服务启动参数
	log.Println("服务启动配置：", srv.Addr)

	// 关闭服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	// 获取停止服务信号，kill  -9 获取不到
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")
	// 执行延迟停止
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting...")
}

func init() {
	// 初始化配置文件
	common.InitConfig()
}

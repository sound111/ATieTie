package main

import (
	"TieTie/dao/myRedis"
	"TieTie/dao/mysql"
	"TieTie/logger"
	"TieTie/pkg/snowflakes"
	"TieTie/routes"
	"TieTie/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title Tietie
// @version 1.0
// @description a simple blog
// @termsOfService http://swagger.io/terms/

// @contact.name sound
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	//加载配置viper
	err := settings.Init("settings/conf.yaml")
	if err != nil {
		fmt.Println("initsettings err:", err)
		return
	}

	//初始化snowflakes
	err = snowflakes.Init()
	if err != nil {
		fmt.Println("initsnowflakes err:", err)
		return
	}

	//初始化日志zap
	err = logger.Init()
	if err != nil {
		fmt.Println("initlogger err:", err)
		return
	}
	defer zap.L().Sync()

	//连接mysql sqlx或gorm
	err = mysql.Init(settings.Conf.MYSQLConfig)
	if err != nil {
		fmt.Println("initmysql err:", err)
		return
	}
	defer mysql.Close()

	fmt.Println(settings.Conf.RedisConfig)
	//连接redis go_redis
	err = myRedis.Init(settings.Conf.RedisConfig)
	if err != nil {
		fmt.Println("initredis err:", err)
		return
	}
	defer myRedis.Close()

	//路由 gin
	gin.SetMode(settings.Conf.AppConfig.Mode)
	r := routes.Setup()

	//启动服务，优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}

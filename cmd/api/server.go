/*
@Time : 2020/11/2 下午9:56
@Author : hoastar
@File : server
@Software: GoLand
*/

package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoastar/orange/database"
	"github.com/hoastar/orange/global/orm"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/hoastar/orange/pkg/task"
	"github.com/hoastar/orange/router"
	"github.com/hoastar/orange/tools"
	config2 "github.com/hoastar/orange/tools/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	config string
	port string
	mode string
	StartCmd = &cobra.Command{
		Use: "server",
		Short: "Start API Server",
		Example: "orange server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")

}

func usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	// 1.读取配置
	config2.ConfigSetup(config)
	// 2.初始化数据库
	database.Setup()
	// 3.启动异步队列
	go task.Start()
}

func run() error {
	if mode != "" {
		config2.SetConfig(config, "settings.application.mode", mode)
	}
	if viper.GetString("settings.application.mode") == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()

	defer func() {
		err := orm.Eloquent.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	if port != "" {
		config2.SetConfig(config, "settings.application", port)
	}

	srv := &http.Server{
		Addr: config2.ApplicationConfig.Host + ":" + config2.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if config2.ApplicationConfig.IsHttps {
			if err := srv.ListenAndServeTLS(config2.SslConfig.Pem, config2.SslConfig.KeyStr); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		}
	}()

	fmt.Printf("%s Server Run http://%s:%s/ \r\n",
		tools.GetCurrentTimeStr(),
		config2.ApplicationConfig.Host,
		config2.ApplicationConfig.Port,
		)

	fmt.Printf("%s Swagger URL http://%s:%s/swagger/index.html \r\n",
		tools.GetCurrentTimeStr(),
		config2.ApplicationConfig.Host,
		config2.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrentTimeStr())

	// 接受中断信号关闭服务器（5秒超时）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown:", err)
	}

	logger.Info("Server exiting")
	return nil
}

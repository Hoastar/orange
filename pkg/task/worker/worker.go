/*
@Time : 2020/11/2 下午10:13
@Author : hoastar
@File : worker
@Software: GoLand
*/

package worker

import (
	"github.com/RichardKnop/machinery/v1"
	taskConfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/hoastar/orange/pkg/logger"
	"github.com/spf13/viper"
)

var AsyncTaskCenter *machinery.Server

func StartServer() {
	tc, err := NewTaskCenter()
	if err != nil {
		panic(err)
	}
	AsyncTaskCenter = tc
}


func NewTaskCenter() (*machinery.Server, error) {
	cnf := &taskConfig.Config{
		Broker: viper.GetString("settings.redis.url"),
		DefaultQueue: "ServerTaskQueue",
		ResultBackend: "eager",
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}
	initAsyncTaskMap()
	return server, server.RegisterTasks(asyncTaskMap)
}

func NewAsyncTaskWorker(concurrency int) *machinery.Worker {
	consumerTag := "TaskWorker"
	worker := AsyncTaskCenter.NewWorker(consumerTag, concurrency)
	errorHandler := func(err error) {
		logger.Error("执行失败: ", err)
	}

	preTaskHandler := func(signature *tasks.Signature) {
		logger.Info("开始执行: ", signature.Name)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		logger.Info("执行结束: ", signature.Name)
	}

	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorHandler)
	worker.SetPreTaskHandler(preTaskHandler)
	return worker
}



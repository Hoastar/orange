/*
@Time : 2020/11/2 下午10:14
@Author : hoastar
@File : server
@Software: GoLand
*/

package task

import (
	"github.com/hoastar/orange/pkg/logger"
	"github.com/hoastar/orange/pkg/task/worker"
)

func Start() {
	// 1. 启动服务，连接redis
	worker.StartServer()

	// 2. 启动异步调度
	taskWorker := worker.NewAsyncTaskWorker(10)
	err := taskWorker.Launch()
	if err != nil {
		logger.Errorf("启动machinery失败，%v", err.Error())
	}
}

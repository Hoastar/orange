/*
@Time : 2020/11/2 下午10:14
@Author : hoastar
@File : send
@Software: GoLand
*/

package task

import (
	"context"
	"github.com/hoastar/orange/pkg/task/worker"
)

func Send(classify string, scriptPath string, params string) {
	worker.SendTask(context.Background(), classify, scriptPath, params)
}
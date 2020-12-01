/*
@Time : 2020/11/28 下午6:20
@Author : hoastar
@File : task
@Software: GoLand
*/

package service

import (
	"fmt"
	"github.com/hoastar/orange/pkg/task"
	"github.com/spf13/viper"
	"strings"
)

func ExecTask(taskList []string, params string) {
	for _, taskName := range taskList {
		filePath := fmt.Sprintf("%v/%v", viper.GetString("script.path"), taskName)
		if strings.HasSuffix(filePath, ".py") {
			task.Send("python", filePath, params)
		} else if strings.HasSuffix(filePath, ".sh") {
			task.Send("shell", filePath, params)
		}
	}
}
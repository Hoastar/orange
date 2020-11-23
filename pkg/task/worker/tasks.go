/*
@Time : 2020/11/2 下午10:13
@Author : hoastar
@File : tasks
@Software: GoLand
*/

package worker

import (
	"context"
	"errors"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/hoastar/orange/pkg/logger"
	"os/exec"
	"syscall"
)

var asyncTaskMap map[string]interface{}

// 任务执行基础
func executeTaskBase(scriptPath string, params string) (err error) {
	// 初始化Cmd
	command := exec.Command(scriptPath, params)
	out, err := command.CombinedOutput()
	if err != nil {
		logger.Errorf("task exec failed, %v", err.Error())
		return
	}
	logger.Info("Output: ", string(out))
	logger.Info("ProcessState PID: ", command.ProcessState.Pid())
	logger.Info("Exit Code ", command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
	return
}

// ExecCommand 异步任务
func ExecCommand(classify string, scriptPath string, params string) (err error) {
	if classify == "shell" {
		logger.Info("start exec shell - ", scriptPath)
		err = executeTaskBase(scriptPath, params)
		if err != nil {
			return
		}
	} else if classify == "python" {
		logger.Info("start exec python - ", scriptPath)
		err = executeTaskBase(scriptPath, params)
		if err != nil {
			return
		}
	} else {
		err = errors.New("当前仅支持python与shell脚本任务，请知悉")
		return
	}
	return
}

func SendTask(ctx context.Context, classify string, scriptPath string, params string) {
	args := make([]tasks.Arg, 0)
	args = append(args, tasks.Arg{
		Name: "classify",
		Type: "string",
		Value: classify,
	})

	args = append(args, tasks.Arg{
		Name:  "scriptPath",
		Type:  "string",
		Value: scriptPath,
	})

	args = append(args, tasks.Arg{
		Name:  "params",
		Type:  "string",
		Value: params,
	})

	task, _ := tasks.NewSignature("ExecCommandTask", args)
	task.RetryCount = 5
	_, err := AsyncTaskCenter.SendTaskWithContext(ctx, task)

	if err != nil {
		logger.Error(err.Error())
	}
}

func initAsyncTaskMap() {
	asyncTaskMap = make(map[string]interface{})
	asyncTaskMap["ExecCommandTask"] = ExecCommand
}




































































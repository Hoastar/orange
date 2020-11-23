/*
@Time : 2020/11/5 下午10:21
@Author : hoastar
@File : utils
@Software: GoLand
*/

package tools

import (
	"github.com/hoastar/orange/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// 校验密码正确性
func CompareHashAndPassword(e, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		logger.Info(err.Error())
		return false, err
	}
	return true, nil
}


// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码

func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}

		if msg == "" {
			msg = err.Error()
		}
		logger.Info(err)
		panic("CustomError" + strconv.Itoa(statusCode) + "#" + msg)
	}
}


func StrToInt(err error, index string) int {
	result, err := strconv.Atoi(index)
	if err != nil {
		HasError(err, "string to int err"+err.Error(), -1)
	}
	return result
}


// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, msg string, code ...int) {
	if !condition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}


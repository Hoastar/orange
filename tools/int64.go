/*
@Time : 2020/11/14 上午12:52
@Author : hoastar
@File : int64
@Software: GoLand
*/

package tools

import "strconv"

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

/*
@Time : 2020/11/3 下午11:25
@Author : hoastar
@File : float64Tostring
@Software: GoLand
*/

package tools

import "strconv"

func FloatToString(e float64) string {
	return strconv.FormatFloat(e, 'E', -1, 64)
}

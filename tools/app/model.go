/*
@Time : 2020/11/6 上午12:02
@Author : hoastar
@File : model
@Software: GoLand
*/

package app

type Response struct {
	// 状态码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

type Page struct {
	List interface{} `json:"list"`
	Count int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize int `json:"pageSize"`
}

type PageResponse struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data Page `json:"data"`
	Msg string `json:"msg"`
}


func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}
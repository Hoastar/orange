/*
@Time : 2020/11/2 下午10:38
@Author : hoastar
@File : string
@Software: GoLand
*/

package tools

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

func GetCurrentTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}


func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err != nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}


func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
			return nil, err
	}
		mapData := make(map[string]interface{})
		err = json.Unmarshal(dataBytes, &mapData)
		if err != nil {
			return nil, err
		}
	return mapData, nil
}




































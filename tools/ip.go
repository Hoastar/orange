/*
@Time : 2020/11/3 下午11:13
@Author : hoastar
@File : ip
@Software: GoLand
*/

package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "局域网ip"
	}

	// 高地地图接口服务：根据ip确定大致地理位置
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))

	m := make (map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		fmt.Println("Umarshal failed", err)
	}

	if m["province"] == "" {
		return "位置位置"
	}
	return m["province"] + "_" + m["city"]
}

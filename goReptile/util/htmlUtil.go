package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHtmlStr(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("GetHtmlStr get err:", err)
		return ""
	}
	// fmt.Println(resp.Header) //头
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetHtmlStr read err:", err)
		return ""
	}
	// 编码转换
	return ConvertToString(string(body), "gbk", "utf-8")
}

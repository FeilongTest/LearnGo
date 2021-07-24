package main

import (
	"github.com/levigross/grequests"
	"log"
	"strings"
)

func main()  {
	session := grequests.NewSession(nil)
	//固定jwtToken
	var submitHeader map[string]string //这里创建临时提交的Header
	submitHeader["Host"] = "feilongwl.cn"
	submitHeader["Content-Type"] = "application/json; charset=utf-8"
	resp, err := session.Post("api.feilongwl.cn",
		&grequests.RequestOptions{
			JSON:    strings.NewReader("{\"data\":\"123\"}"),
			Headers: submitHeader,
		})
	if err != nil {
		log.Printf("发送失败：%s\n", err)
	} else {
		log.Printf("发送成功：%s\n", resp.Bytes())
	}
}

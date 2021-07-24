package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)


type WssData struct {
	Client      string `json:"client"`
	Module      string `json:"module"`
	ModuleGroup string `json:"moduleGroup"`
	Url         string `json:"url"`
	Duration    int    `json:"duration"`
	Concurrent  int    `json:"concurrent"`
	TTL         int    `json:"TTL"`
}

func sendMessage(ws *websocket.Conn,data []byte) {
	ws.Write(data)
	readMessage(ws)
}

func readMessage(ws *websocket.Conn) {
	data := make([]byte, 1024*10)
	m, err := ws.Read(data)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("读取数据：",string(data[:m]))
}


func main()  {


	openid := "74b90d41ae584a159eff94513160c499"
	jwtToken := "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJlZHgudW5pcHVzLmNuIiwiYWRtaW5pc3RyYXRvciI6ZmFsc2UsIm9wZW5faWQiOiI3NGI5MGQ0MWFlNTg0YTE1OWVmZjk0NTEzMTYwYzQ5OSIsIm5hbWUiOiIiLCJpc3MiOiJjNGY3NzIwNjNkY2ZhOThlOWM1MCIsImVtYWlsIjoiIiwiZXhwIjoxNjU3NjEyNDM0fQ.HseqVEHlR39HC5_WyUzAwfljpYqELVHHj5AHW2O2f-s"
	courseUrl := "u1g20"
	courseId := "course-v1:Unipus+ctew01+2018_03"
	pageUrl := " "//假设为空

	submitActivity := fmt.Sprintf("40/userActivities?uuid=%s&token=%s",openid,jwtToken)

	submitUrl := fmt.Sprintf("wss://ucontent.unipus.cn/unipusio/?EIO=3&transport=websocket&uuid=%s&token=%s",openid, jwtToken)

	ws, err := websocket.Dial(submitUrl, "", "https://ucontent.unipus.cn/")
	if err != nil{
		fmt.Println("发送websocket出错",err)
	}
	readMessage(ws)

	sendMessage(ws,[]byte(submitActivity))


	kickData := WssData{
		Client:      "U校园mobile",
		Module:      courseUrl,
		ModuleGroup: courseId,
		Url:         pageUrl,
		Duration:    20,//区域内随机，时长参数
		Concurrent:  1,
		TTL:         60,
	}

	//结构体转json
	value, err := json.Marshal(&kickData)
	if err != nil{
		fmt.Println("wss数据转换json出错",err)
	}

	wssData := "42/userActivities,0[\"kick\","+ string(value) +"]"

	sendMessage(ws,[]byte(wssData))



	defer ws.Close()
	println("main over")

}



//检查错误
func checkError(err error) {
	if err != nil {
		fmt.Errorf("检测到错误! %s", err.Error)
	}
}

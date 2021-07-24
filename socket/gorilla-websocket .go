package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
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

func main() {
	openid := "74b90d41ae584a159eff94513160c499"
	jwtToken := "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJlZHgudW5pcHVzLmNuIiwiYWRtaW5pc3RyYXRvciI6ZmFsc2UsIm9wZW5faWQiOiI3NGI5MGQ0MWFlNTg0YTE1OWVmZjk0NTEzMTYwYzQ5OSIsIm5hbWUiOiIiLCJpc3MiOiJjNGY3NzIwNjNkY2ZhOThlOWM1MCIsImVtYWlsIjoiIiwiZXhwIjoxNjU3NjEyNDM0fQ.HseqVEHlR39HC5_WyUzAwfljpYqELVHHj5AHW2O2f-s"
	courseUrl := "u0g1"
	courseId := "course-v1:Unipus+ctew01+2018_03"
	pageUrl := " "//假设为空


	submitUrl := "EIO=3&transport=websocket&uuid="+openid+"&token="+jwtToken
	submitActivity := "40/userActivities?uuid="+openid+"&token="+jwtToken


	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{
		Scheme: "wss",
		Host: "ucontent.unipus.cn",
		Path: "/unipusio/",
		RawQuery: submitUrl,
	}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), http.Header{
		"Origin": 				  []string{"https://ucontent.unipus.cn/"},
		"User-Agent": 		 	  []string{"okhttp/3.8.1"},
	})

	if err != nil {
		log.Fatal("创建socket请求出错:", err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if message != nil {
				log.Printf("接受到消息: %s", message)
				if find := strings.Contains(string(message), "\"msg\":\"ok\","); find {
					//done <- struct{}{}
					break

				}
			}

		}
	}()

	//创建完成延迟1s
	//time.Sleep(time.Second)






	err = c.WriteMessage(websocket.TextMessage, []byte(submitActivity))
	if err != nil {
		log.Println("write:", err)
		return
	}
	//发送完成延迟1s
	time.Sleep(time.Second)

	kickData := WssData{
		Client:      "U校园mobile",
		Module:      courseUrl,
		ModuleGroup: courseId,
		Url:         pageUrl,
		Duration:    1,//区域内随机，时长参数
		Concurrent:  1,
		TTL:         60,
	}

	//结构体转json
	value, err := json.Marshal(&kickData)
	if err != nil{
		fmt.Println("wss数据转换json出错",err)
	}

	wssData := "42/userActivities,0[\"kick\","+ string(value) +"]"

	err = c.WriteMessage(websocket.TextMessage, []byte(wssData))
	if err != nil {
		log.Println("write:", err)
		return
	}

	defer c.Close()



	for {
		select {
		case <-done:
			log.Println("main over")
			return
		}
	}
}


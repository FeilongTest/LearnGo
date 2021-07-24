package main

import (
	"fmt"
	"sync"
	"time"
)


/*
 * 分别使用Ticker和Timer创建耗时协程A,B,C
 * 将A,B,C三个协程加入等待组
 * A,B,C结束时从等待组注销
 * 主协程阻塞至等待组内协程数归零
*/

func main()  {

	//声明一个等待组
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		fmt.Println("协程A开始工作")
		time.Sleep(3*time.Second)
		fmt.Println("协程A 结束")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程B开始工作")
		<-time.After(5*time.Second)
		fmt.Println("协程B 结束")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程C开始工作")
		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 4; i++ {
			<-ticker.C
		}
		ticker.Stop()
		fmt.Println("协程C 结束")
		wg.Done()

	}()

	//阻塞等待wg中的协程数量归零
	wg.Wait()
	fmt.Println("主协程 结束")

}

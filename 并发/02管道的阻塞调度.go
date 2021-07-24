package main

import (
	"fmt"
	"time"
)

func Count(grName string, n int,chanQuit chan string)  {
	for i := 0; i <= n; i++ {
		fmt.Println(grName,i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println(grName,"工作完毕")
	chanQuit <- grName+"over"
	//2.向通知管道中写入

}

/*
 * 如果没有管道堵塞 会在子协程执行不完的情况结束主协程
 * 所以直接加入管道阻塞机制
 */
func main()  {
	//1.创建一个3缓存的管道

	chanQuit := make(chan string,3)


	go Count("son",10,chanQuit)
	go Count("daughter",7,chanQuit)
	Count("main",5,chanQuit)

	//3.阻塞等待从【任务完毕通知管道】读出所有协程的结束消息
	for i := 0; i < 3; i++ {
		x := <- chanQuit
		fmt.Println(x)
	}

	//恰好结束在所有协程执行完毕时
	println("main over")


}
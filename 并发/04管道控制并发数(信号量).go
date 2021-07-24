package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

/*
 * 100条写成并发求1-10000平方根
 * 最大并发数控制在5
 * 管道实现
*/

func GetSqrt(name string, n int,chSem chan string) {
	//想执行协程 必须先注册
	//能写入就能执行 写不进去就阻塞到能写入为止 （管道的特性）
	chSem <- name
	ret := math.Sqrt(float64(n))
	time.Sleep(time.Second)
	fmt.Printf("%d的平方根是：%.2f\n",n,ret)

	//任务执行完毕 从信号量控制管道注销自己 以便为其他并发任务腾出空间
	<-chSem
}


func main()  {


	//并发数（信号量）控制管道 凡是要并发执行的协程必须先将协程名字注册到该管道
	var chSem = make(chan string,5)


	for i := 0; i < 100; i++ {
		go GetSqrt("协程"+strconv.Itoa(i),i,chSem)
	}

	//永远等待
	for  {
		time.Sleep(time.Second)
	}
}

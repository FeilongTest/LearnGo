package main

import "fmt"

func main()  {

	var myChan chan int
	fmt.Println(myChan)

	//创建缓存能力为0的管道
	myChan = make(chan int)
	fmt.Println(myChan)

	//创建缓存能力为1的管道
	myChan = make(chan int,1)
	fmt.Println(myChan)

	myChan <- 123
	fmt.Println("main over")



}

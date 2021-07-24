package main

/*
 * 一写 两读 三条管道 随机选择一条能走的路
 * 等所有路都走不通 退出循环
*/

func main()  {

	chA := make(chan int,5)
	chB := make(chan int,4)
	chC := make(chan int,3)

	go TaskA(chA)
	go TaskB(chB)
	go TaskC(chC)



}

func TaskA(ch chan int) {
	ch <- 123
}

func TaskB(ch chan int) {
	<-ch
}

func TaskC(ch chan int) {
	<-ch
}
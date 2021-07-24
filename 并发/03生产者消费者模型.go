package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
 * 生产者每天生产一件商品
 * 生产者每生产一件产品，消费者就消费一件
 * 消费10轮后，主协程退出
*/

/*
 * 思路：
 * 生产者源源不断向【商店管道】写入产品
 * 消费者阻塞等待从【商店管道】读出数据
 * 消费者每读出一次数据（消费一次）。就向【计数管道】写入一个数据
 * 主协程阻塞从【计数管道】读出10个数据就结束
*/

type Product struct {
	Name string
}


func main()  {

	//创建商店管道、计数管道
	chanShop  := make(chan Product,5)
	chanCount := make(chan int,5)

	//生产者源源不断向【商店管道】写入产品
	go Produce(chanShop)
	//消费者阻塞等待从【商店管道】读出数据
	//消费者每读出一次数据（消费一次）。就向【计数管道】写入一个数据
	go Consume(chanShop,chanCount)

	//主协程阻塞从【计数管道】读出10个数据就结束
	for i := 0; i < 10; i++ {
		<-chanCount
	}
	fmt.Println("main over")
}

/*
 *chanShop <- chan Product 此时的chanShop是只读管道 receive-only channel
 *chanCount chan <-int 此时的chanCount是只写管道 send-only channel
*/

func Consume(chanShop <-chan Product,chanCount chan <-int) {
	for {
		product := <-chanShop
		fmt.Println("消费者消费了",product)

		//向计数管道写入数据
		chanCount <- 1
	}
}

/*
 *chanShop chan<-Product 此时的chanShop是只写管道 send-only channel
 */

func Produce(chanShop chan<-Product) {
	for {
		//生产商品
		product := Product{Name: "产品"+strconv.Itoa(time.Now().Second())}
		//丢入商店
		chanShop <- product
		fmt.Println("生产者向商店输送了",product)
		//等待一天
		time.Sleep(1 * time.Second)
	}
}

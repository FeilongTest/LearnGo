package main

import (
	"fmt"
	"sync"
)

/* 使用10个协程他同时频繁修改一个数据 造成并发不安全的问题*/

func main()  {

	var money = 2000
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				money += 1
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("钱数=",money)
}
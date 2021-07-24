package main

import (
	"fmt"
	"time"
)

func main()  {
	//tiker := time.NewTicker(15 * time.Minute)
	tiker := time.NewTicker(1 * time.Second)

	var count int
	for {
		//go func() {
		//	fmt.Println("计数",count)
		//}()
		fmt.Println("计数",count)
		count++
		<-tiker.C
	}
}


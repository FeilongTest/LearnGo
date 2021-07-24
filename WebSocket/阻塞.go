package main

import "fmt"

func main()  {

	chanQuit := make(chan bool,1)

	go func() {
		for i := 0; i < 10; i++ {
			if i == 9{
				chanQuit <- true
			}
			fmt.Println(i)
		}
	}()







	for {
		_, ok := <-chanQuit
		if ok {
			break
		}
	}
	fmt.Println("main over")

}
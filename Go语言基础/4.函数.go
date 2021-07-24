package main

import "fmt"

func main(){

}

//闭包 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。
func closureAdder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func closure(){
	var f = closureAdder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := closureAdder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}

//defer


//panic/recover

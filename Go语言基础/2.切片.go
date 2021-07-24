package main

import "fmt"

func main() {
	funSlice1()
}
//定义切片
func funSlice1(){
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}


//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
func funSlice2(){
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println("切片内容为：",s1,"切片长度为：",len(s1),"切片深度为：",cap(s1),"切片是否为nil：",s1==nil)
	fmt.Println("切片内容为：",s2,"切片长度为：",len(s2),"切片深度为：",cap(s2),"切片是否为nil：",s2==nil)
	fmt.Println("切片内容为：",s3,"切片长度为：",len(s3),"切片深度为：",cap(s3),"切片是否为nil：",s3==nil)

	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0{
		fmt.Println("切片为空")
	}
}

//切片的赋值拷贝 下面的代码拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

func funSliceCopy() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
	//	copy(destSlice, srcSlice []T)也可以实现同样的功能
	//    srcSlice: 数据来源切片
	//    destSlice: 目标切片
}

// 切片删除元素
func funSliceDelete() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
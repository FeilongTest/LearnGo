package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//打印类型
func 打印变量类型(){
	var i int
	log.Println(reflect.TypeOf(i))
}

func 时间字符串互转1()  {
	temp := "2021.03.01-2021.07.30"
	pos := strings.Index(temp, "-")
	temp = temp[pos+1:]
	//temp  = strings.Replace(temp, ".", "-", -1)
	//layout为golang默认模版
	stamp, _ := time.ParseInLocation("2006.01.02", temp, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())  //输出：1546926630
	//log.Println()
	if stamp.Unix() < time.Now().Unix() {
		log.Println("过期了")
	}

}

func 时间字符串互转(){
	t := int64(1546926630)      //外部传入的时间戳（秒为单位），必须为int64类型
	t1 := "2019-01-08 13:50:30" //外部传入的时间字符串

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	timeTemplate2 := "2006/01/02 15:04:05" //其他类型
	timeTemplate3 := "2006-01-02"          //其他类型
	timeTemplate4 := "15:04:05"            //其他类型

	// ======= 将时间戳格式化为日期字符串 =======
	log.Println(time.Unix(t, 0).Format(timeTemplate1)) //输出：2019-01-08 13:50:30
	log.Println(time.Unix(t, 0).Format(timeTemplate2)) //输出：2019/01/08 13:50:30
	log.Println(time.Unix(t, 0).Format(timeTemplate3)) //输出：2019-01-08
	log.Println(time.Unix(t, 0).Format(timeTemplate4)) //输出：13:50:30

	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())  //输出：1546926630
}

func main转换001()  {
	//string与int互转
	var num1 int = 10;
	//Itoa底层调用的是FormatInt
	//I to S
	str1 := strconv.Itoa(num1)
	fmt.Println(str1)
	//S to I
	num1_int, _ := strconv.Atoi(str1)
	fmt.Println(num1_int)

	//int64与string类型
	var num2 int64 = 432;
	//I to S
	str2 := strconv.FormatInt(num2, 10)
	fmt.Println(str2)
	//S to I
	num2_int, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Println(num2_int)

	//float与string互转
	//bitSize表示最后一位的位数设置为float32或者float64类型
	var f1 float64 = 12.432
	//F to S
	str3 := strconv.FormatFloat(f1, 'E', -1, 32)
	fmt.Println(str3)
	//S to F
	f_float, _ := strconv.ParseFloat(str3, 32)
	fmt.Println(f_float)

	//	bool与string互转
	var bb bool = true
	//B to S
	str4 := strconv.FormatBool(bb)
	fmt.Println(str4)
	//S to B
	b, _ := strconv.ParseBool(str4)
	fmt.Println(b)

	//interface转其他类型————返回值是interface，直接赋值是无法转化的
	//interface 转string
	var a interface{}
	var str5 string
	a = "3432423"
	str5 = a.(string)
	fmt.Println(str5)

	//interface 转int
	var m interface{}
	var m1 int
	m = 43
	m1 = m.(int)
	fmt.Println(m1)

	//interface 转float64
	var ff interface{}
	var ff1 float64
	ff = 432.54
	ff1 = ff.(float64)
	fmt.Println(ff1)
}

package main

import (
	"log"
	"strings"
)

func main()  {
	tracer := "u0g2-103"
	comma := strings.Index(tracer, "-")
	//comma的意思是从字符串tracer查找第一个逗号，然后返回他的位置，这里的每个中文是占3个字符，从0开始计算，那么逗号的位置就是12
	log.Println(tracer[comma+1:])
}


//字符串截取
func stringIntercept()  {
	tracer := "feilong|test"
	comma := strings.Index(tracer, "|")
	//comma的意思是从字符串tracer查找第一个逗号，然后返回他的位置，这里的每个中文是占3个字符，从0开始计算，那么逗号的位置就是12
	log.Println(tracer[:comma])
}

//字符串替换
func stringReplace()  {
	//-1 为所有 指定n则为替换n次
	stringTest := `{\n   \"questionNumber\": \"25\",\n   \"exerciseType\": 3,\n   \"fromApp\": 3001,\n   \"tutorialId\": \"course-v1:Unipus+nhce_3_vls_1+2016_10\",\n   \"totalScore\": \"25\",\n   \"exerciseName\": \"Unit test\",\n   \"scoreDetail\": [\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1\n   ],\n   \"type\": \"ut\",\n   \"tutorialType\": 2,\n   \"exerciseId\": \"454\",\n   \"examId\": \"20000083\"\n}`
	stringTest  = strings.Replace("stringTest", "\\n", "", -1)
	stringTest  = strings.Replace(stringTest, " ", "", -1)
	stringTest  = strings.Replace(stringTest, "\\", "", -1)
}

//字符串包含
func stringContains()  {
	if strings.Contains("abc","a"){
		log.Println("包含")
	}
}
package main

import (
	"encoding/json"
	"log"
	"strings"
)



//处理不确定json
func uncertainJson()  {
	stringTest := `{\n   \"questionNumber\": \"25\",\n   \"exerciseType\": 3,\n   \"fromApp\": 3001,\n   \"tutorialId\": \"course-v1:Unipus+nhce_3_vls_1+2016_10\",\n   \"totalScore\": \"25\",\n   \"exerciseName\": \"Unit test\",\n   \"scoreDetail\": [\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1,\n      1\n   ],\n   \"type\": \"ut\",\n   \"tutorialType\": 2,\n   \"exerciseId\": \"454\",\n   \"examId\": \"20000083\"\n}`
	stringTest  = strings.Replace(stringTest, "\\n", "", -1)
	//stringTest  = strings.Replace(stringTest, " ", "", -1)
	stringTest  = strings.Replace(stringTest, "\\", "", -1)
	log.Println(stringTest)
	var summary map[string]json.RawMessage
	if err := json.Unmarshal([]byte(stringTest),&summary); err != nil{
		log.Println("ut中summary参数 转换json出错")
	}
	log.Println(string(summary["exerciseName"]))
}


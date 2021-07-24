package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {
	//13位时间戳
	fmt.Println(strconv.FormatInt(time.Now().UnixNano() / 1e6,10))


}
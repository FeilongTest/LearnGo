package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

//密钥 key
const (
	SIGNED_KEY = "7d8b0b1a8e2746ec9b32eddd1aa9acc2"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	OpenId        string `json:"open_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Administrator bool   `json:"administrator"`
	jwt.StandardClaims
}

//加密
func jwtTokenEncode() string{
	/*
	//另外一种声明claims方式
	claim := jwt.MapClaims{
		"open_id":"a2a7fd8138ab480ea96a6ccfcb783a00",//固定
		"name": "",
		"email": "",
		"administrator": false,
		"exp": time.Now().UnixNano() / 1e6,//13位时间戳
		"iss": "c4f772063dcfa98e9c50",//固定
		"aud": "edx.unipus.cn",//固定
	}
*/
	clims := CustomClaims{
		OpenId:"a2a7fd8138ab480ea96a6ccfcb783a00",//固定
		Name: "",
		Email: "",
		Administrator: false,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().UnixNano() / 1e6,//13位时间戳
			Issuer: "c4f772063dcfa98e9c50",//固定
			Audience: "edx.unipus.cn",//固定
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clims)
	ss, err := token.SignedString([]byte(SIGNED_KEY))
	if err != nil {
		log.Println(err)
	}
	return  ss
}



func main()  {
	fmt.Println(jwtTokenEncode())
}
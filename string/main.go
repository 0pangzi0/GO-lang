package main

import (
	"fmt"
	"strings"
)

//字符串

func main() {
	// \ 本来是具有特殊含义的,我应该告诉程序我写的\就是一个单纯的\
	//path := "\"D:\\GO\\src\\code/oldboyedu.com\\studygo\\day01"\"
	//fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)
	//多行的字符串
	s2 := `
		大雨磅礴
		百鬼夜行
		有些人混在里面
		比鬼还开心
	`
	fmt.Println(s2)
	s3 := `/Users/w-w/go/go/string`
	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "傻逼"

	ss := name + world
	//fmt.Println(name + world)
	fmt.Println(ss)
	//fmt.Printf("%s%s\n", name, world)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	//分割  以 / 为分割
	ret := strings.Split(s3, "/")
	fmt.Println(ret)

	//包含  判断是否包含理想或者理性
	fmt.Printf("判断是否包含理性:")
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Printf("判断是否包含理想:")
	fmt.Println(strings.Contains(ss, "理想"))

	//前缀  判断前缀是不是  理想
	fmt.Printf("判断前缀是不是理想:")
	fmt.Println(strings.HasPrefix(ss, "理想"))
	//后缀  判断后缀是不是  理想
	fmt.Printf("判断后缀是不是理想:")
	fmt.Println(strings.HasSuffix(ss, "理想"))

	//子串出现的位置
	s4 := "abcdeb"
	fmt.Printf("判断c出现的位置:")
	fmt.Println(strings.Index(s4, "c"))
	fmt.Printf("判断b最后一次出现的位置:")
	fmt.Println(strings.LastIndex(s4, "b"))

	//拼接  使用+把ret连接起来
	fmt.Println(strings.Join(ret, "+"))
}

package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n) //查看类型    打印值的类型
	fmt.Printf("%v\n", n) // :值   值的默认格式表示  不知道什么类型用  %v 是通用的
	fmt.Printf("%b\n", n) //:二进制数   整型 : 表示为二进制  :浮点数与复数: 无小数部分,二进制指数的科学计数法,如-123456p-78
	fmt.Printf("%d\n", n) //:十进制数   整型 : 表示为十进制
	fmt.Printf("%o\n", n) //:八进制数   整型 : 表示为八进制
	fmt.Printf("%x\n", n) //:十六进制数 整型 : 表示为十六进制,使用a-f
	var s = "hello,世界!"
	fmt.Printf("字符串:%s\n", s)  //:字符串    直接输出字符串或者  []byte
	fmt.Printf("字符串:%v\n", s)  // :值      值的默认格式表示    不知道什么类型用  %v 是通用的
	fmt.Printf("字符串:%#v\n", s) //:值的GO语法表示
}

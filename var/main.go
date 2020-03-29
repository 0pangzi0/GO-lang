package main

import "fmt"

//Go语言中推荐使用驼峰式命名
//声明变量
//var name string
//var age int
//var isOk bool

//批量声明
var (
	name string //""
	age  int    //0
	isOk bool   //false
)

func main() {
	name = "理想"
	age = 24
	isOk = true
	//Go语言中变量声明必须使用,不使用就编译不过去
	fmt.Print(isOk)             //在终端中输出要打印的内容
	fmt.Printf("name:%s", name) //%s占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)            //打印完指定内容之后会在后面加一个换行符
	fmt.Println(name, age, isOk)
	// 声明变量同时赋值
	var s1 string = "皮小妹"
	fmt.Println(s1)
	//类型推导(根据值判断该变量是什么类型)
	var s2 = "22"
	fmt.Println(s2)
}

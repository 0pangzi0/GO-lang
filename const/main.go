package main

import "fmt"

//常量
//定义了常量之后不能更改
//在程序运行期间不会改变的量
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

//批量声明常量时,如果某一行声明后没有赋值的,默认就和上一行一样
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2 = iota //1
	_  = iota //2
	b3 = iota //3
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2 口吐芬芳 在一个const里面不会重置为0
	c4        //3
)

//多个常量声明在一行,const中每新增一行常量声明将使iota计数一次
const (
	d1, d2 = iota + 1, iota + 2 //d1:1 d2:2  0+1  0+2 const中每新增一行常量声明将使iota计数一次
	d3, d4 = iota + 1, iota + 2 //d3:2 d4:3  1+1  1+2 const中每新增一行常量声明将使iota计数一次
)

//定义数量级
const(
	_ = iota  //把iota扔了,目前不需要它
	KB = 1 << (10* iota)  //<<像左移十位
	MB = 1 << (10* iota)
	GB = 1 << (10* iota)
	TB = 1 << (10* iota)
	PB = 1 << (10* iota)
)

func main() {
	//pi = 123
	//fmt.Println("pi:", pi)

	//fmt.Println("statusOK", statusOK)

	//fmt.Println("notFound:", notFound)

	//fmt.Println("n1:", n1)
	//fmt.Println("n2:", n2)
	//fmt.Println("n3:", n3)

	//fmt.Println("a1:", a1)
	//fmt.Println("a2:", a2)
	//fmt.Println("a3:", a3)

	//fmt.Println("b1:", b1)
	//fmt.Println("b2:", b2)
	//fmt.Println("b3:", b3)

	//fmt.Println("c1:", c1)
	//fmt.Println("c2:", c2)
	//fmt.Println("c3:", c3)
	//fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
}

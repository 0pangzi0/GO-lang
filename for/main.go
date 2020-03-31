package main

import "fmt"

//for循环

func main() {
	//基本格式
	// for i := 0; i < 10; i++ { //i等于0,i小于10,i加加在原来的基础上加一
	// 	fmt.Println(i)
	// }

	//变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	//变种2
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//无限循环
	// for {
	// 	fmt.Println("123")
	// }

	//for range循环
	s := "hello中国"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

}

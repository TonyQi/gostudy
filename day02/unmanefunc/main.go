package main

import "fmt"

//匿名函数
func main() {
	//在函数内部无法声明另一个代名字的函数
	var f1 = func(x, y int) int {
		return x + y
	}
	ret := f1(1, 2)
	fmt.Printf("ret is %d ;\n", ret)
	//立即执行函数,只执行一次，调用完就销毁
	ret2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Printf("ret2 is %v \n", ret2)

}

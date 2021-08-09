package main

import (
	"fmt"
)

//高阶函数
//函数类型

func f1() {
	fmt.Println("1111")
}
func f2(x string) int {
	fmt.Println(x)
	return 1
}

//函数类型作为输入值

func f3(x func(string) int) {
	ret := x("111")
	fmt.Printf("ret is %v \n ", ret)
}

//函数还可以作为返回值
func f4(x func()) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("a is %T \n", a)
	b := f2
	fmt.Printf("b is %T \n", b)
	f3(b)

}

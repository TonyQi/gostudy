package main

import (
	"fmt"
)

//闭包是一个函数，这个函数有个特点，包含了一个外部函数的变量，
//底层原理，函数作为返回值，函数查找变量顺序，内部找，再外部找
func f1(f func()) {
	f()
}

func f2(x, y int) {
	fmt.Printf("%d + %d is %d \n", x, y, x+y)
}

func f3(x func(int, int), m, n int) func() {
	return func() {
		x(m, n)
	}
}

func main() {
	//要把f2传入f1,需要对f2进行包装，函数见f3
	f1(f3(f2, 1, 2))
}

package main

import "fmt"

//基于基本类型与类型别名
//自定义类型
type myInt int

//类型别名
type yourInt = int

type person struct {
	name string
	age  int
	sex  bool
}

func main() {
	var n myInt
	n = 100
	fmt.Printf("n is %T \n", n)
	var m yourInt
	m = 100
	fmt.Printf("m is %T \n", m)

}

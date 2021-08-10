package main

import "fmt"

//构造体构造函数
type person struct {
	name string
	age  int
}

//返回结构体还是指针要有考虑
//当结构体比较大的时候，返回指针，结构体比较小的时候使用结构体
func newPersonFunc(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPersonFunc("123", 12)
	fmt.Println(p1)
}

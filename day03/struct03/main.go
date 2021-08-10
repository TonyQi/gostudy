package main

import "fmt"

type person struct {
	name string
	age  int
}

//结构体占用一块连续的内存
func main() {
	p1 := person{
		name: "1",
		age:  1,
	}
	fmt.Printf("name is %v \n", &(p1.name))
	fmt.Printf("name is %v \n", &(p1.age))
}

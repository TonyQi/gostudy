package main

import "fmt"

type person struct {
	name   string
	age    int
	gender bool
	hobby  []string
}

func main() {
	//声明一个person类型的变量
	var p1 person
	//属性赋值
	p1.name = "123"
	p1.age = 18
	p1.gender = false
	p1.hobby = []string{"1", "2", "3"}
	//访问值
	fmt.Printf("p1 is %T \n", p1)
	fmt.Println(p1.name)
	//匿名结构体,多用于临时场景
	var s struct {
		name string
		age  int
	}
	fmt.Printf("s is %T\n ", s)
	//结构体是值类型，不是引用类型,如果要传对象要传指针，具体程序看struct02
	//go中的传参永远是拷贝，是个新的，如果要传引用，需要传指针，
	var p2 person
	p2.name = "123"
	p3 := p2
	p3.age = 123
	fmt.Println(p2, p3)
}

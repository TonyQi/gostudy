package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p1 person
	p1.age = 18
	p1.name = "123"
	p2 := &p1
	p2.name = "123"
	fmt.Println(p1, p2)
	p3 := new(person)
	fmt.Println(p3)
	p3.name = "123"
	p3.age = 12
	//拿到变量对应的内存地址
	fmt.Printf("%v \n", &p3)
	//key value初始化，通常用这种
	p4 := person{
		name: "123",
		age:  12,
	}
	fmt.Println(p4)
	//值的列表初始化,值的顺序必须与定义顺序一致
	p5 := person{
		"1234",
		12,
	}
	fmt.Println(p5)
}

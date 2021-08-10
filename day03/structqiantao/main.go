package main

import "fmt"

//嵌套结构体
type person struct {
	name     string
	age      int
	students []student
}

type student struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name:     "12345",
		age:      12,
		students: []student{student{"1", 1}, student{"2", 2}, student{"3", 3}},
	}
	fmt.Println(p1)
	fmt.Println(p1.students[0].name)
}

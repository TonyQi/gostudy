package main

import "fmt"

//Go语言中首字母大写表示包外部可见
//Peron 这是一个结构体
type Person struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

//方法是针对特定对象的方法

func (d dog) wang() {
	fmt.Println(d.name + "汪汪汪")
}

func (p Person) jiaguonian() {
	p.age = p.age + 1
}

func (p *Person) zhengguonian() {
	p.age = p.age + 1
}

func main() {
	dog1 := new(dog)
	dog1.name = "123"
	dog1.wang()
	p1 := Person{
		name: "123",
		age:  18,
	}
	//注意：在go中所有的传递都是值传递，传递进去是个新对象
	p1.jiaguonian() //age加1是新的对象加1
	fmt.Println(p1.age)
	p1.zhengguonian() //指针传递是传入本对象的指针地址，所以能够加1
	fmt.Println(p1.age)
}

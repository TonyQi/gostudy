package main

import "fmt"

func main() {
	deferdemo()
}

func deferdemo() {
	fmt.Println("1111")
	//一般用于结束前关闭资源操作，例如关闭文件句柄，管理数据库链接等操作，多个defer 按照defer的定义顺序执行
	defer fmt.Println("1123333")
	defer fmt.Println("hahahhaha")
	fmt.Println("3333")
}

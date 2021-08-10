package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	//刚打开数据库链接，需要释放数据库资源
	defer func() {
		error := recover()
		if error != nil {

		}
		fmt.Println("释放数据库")
	}()

	panic("程序崩溃退出")
	fmt.Println("B")
}
func funcC() {
	fmt.Println("C")
}

func main() {

	funcA()
	funcB()
	funcC()
}

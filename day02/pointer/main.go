package main

import "fmt"

func main() {
	//& 取地址   *根据地址取值
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)
}

package main

import "fmt"

const PI = 3.1415964

//批量声明
const (
	URL_NOT_FOUND_ERROR_CODE = 404
	SUCCESS_CODE             = 200
)

//
const (
	a1 = iota //0
	a2        //1
	a3        //2
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(SUCCESS_CODE)
	fmt.Println(PI)
	fmt.Printf("a2", a2)
}

package main

import "fmt"

var name string
var age int
var isOK bool

var (
	fullname   string
	bloodpress float32
)

func main() {
	name = "理想"
	fullname = "lixiang123"
	age = 11
	isOK = false
	bloodpress = 12.2
	fmt.Printf("name is %s", name)

	s3 := 3.141592657
	fmt.Println(s3)
}

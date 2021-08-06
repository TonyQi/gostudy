package main

import "fmt"

func main() {
	i1 := 101
	fmt.Println(i1)
	i2 := 077
	fmt.Printf("i2 is %d ", i2)
	fmt.Printf("i2 is %x", i2)
	fmt.Printf("i2 is %b", i2)

	fmt.Printf("i2 is %T", i2)

	i3 := int8(9)
	fmt.Printf("i3 is %T", i3)
}

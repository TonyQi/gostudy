package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		switch i {
		case 1:
			fmt.Print(i)
			fallthrough
		case 2:
			fmt.Print(i)
		case 3:
			fmt.Print(i)
		case 4:
			fmt.Print(i)
		case 5:
			fmt.Print(i)
		default:
			fmt.Print(i)
		}
	}
}

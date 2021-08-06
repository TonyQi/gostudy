package main

import "fmt"

func main() {
	imax := 9
	for i := 1; i <= imax; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", i, j, i*j)
		}
		fmt.Println("")
	}

}

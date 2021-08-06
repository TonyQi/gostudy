package main

import "fmt"

func main() {
	count := 100

	for i := 0; i < count; i++ {

		if i%5 == 0 {
			//break
			continue
		}
		fmt.Print(i)
	}
}

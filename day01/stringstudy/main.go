package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "善的生态,这同时是Java的优势,也是Java编程语言的一个沉重包袱"
	fmt.Printf("The length of s1 is %d \n", len(s1))
	s1 = strings.ToUpper(s1)
	fmt.Println(s1)
	s2 := []rune(s1)
	s2[2] = '1'
	count := 0
	fmt.Println(string(s2))
	for _, v := range s2 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Printf("The length of Hanzi String is %d \n", count)

}

package main

import (
	"fmt"
	"strings"
)

func addSuffixFunc(suffix string) func(string) string {
	return func(par string) string {
		if !strings.HasSuffix(par, suffix) {
			return par + suffix
		}
		return par
	}
}

func main() {
	txtSuffixFunc := addSuffixFunc(".txt")
	jpgSuffixFunc := addSuffixFunc(".jpg")
	fmt.Println(txtSuffixFunc("test"))
	fmt.Println(jpgSuffixFunc("test"))
}

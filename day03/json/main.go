package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "123",
		Age:  12,
	}
	str, error := json.Marshal(p1)
	if error != nil {
		fmt.Println("Marshal failed!!!")
		return
	}
	fmt.Println(string(str))
	jsons := `{"name":"123456","age":18}`
	p := new(person)
	json.Unmarshal([]byte(jsons), p)
	fmt.Println(*p)

}

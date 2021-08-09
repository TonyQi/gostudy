package main

import "fmt"

//slicen是一个引用类型

func main() {
	//定义一个int类型的切片
	var s []int
	var s2 []string
	fmt.Printf("s is %T,s2 is %T \n", s, s2)
	fmt.Println(s == nil) //true
	//切片初始化
	s = []int{1, 2, 3}
	s2 = []string{"1", "2", "3"}
	// 从数组得到切片
	var array1 = [...]int{1, 2, 3, 5, 6, 7, 8, 9}
	s3 := array1[0:6] //前闭后开
	fmt.Println(s3)
	s3 = array1[:5] //从开始切到5-1
	fmt.Println(s3)
	s3 = array1[3:] //从底第3切到最后
	fmt.Println(s3)
	s3 = array1[:] //从头切到尾
	fmt.Println(s3)
	fmt.Printf("len is %d, cap is %d \n", len(s3), cap(s3))
	//切片容量是指从切片的第一个元素到最后的长度
	//切片再切片
	s4 := s3[3:]
	fmt.Println(s4)
}

package main

import "fmt"

func main() {
	//数组定义
	//数组必须指定存放类型和存放长度
	//数组长度和类型数组类型的一部分，所以a1和a2无法比较
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1 is %T \n", a1)
	fmt.Printf("a2 is %T \n", a2)
	//数组初始化，数组不初始化，默认都是0值
	//数组初始化方案1
	a1 = [3]bool{true, true, true}
	//数组初始化方案2,根据初始值自动推断长度
	a100 := [...]int{0, 1, 1, 2, 23, 5, 8, 6, 99, 4, 5, 65, 85, 66, 95, 5, 55, 5, 5, 5, 5, 5, 565266354}
	fmt.Println(a100)
	//数组初始化方案3，根据索引指定值
	a3 := [5]int{0: 1, 4: 12}
	fmt.Println(a3)

	//根据索引遍历
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}

	// for i, v := range a100 {
	// 	fmt.Println(i, v)
	// }

	//多维数组
	var a4 [3][2]int
	a4 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a4)
	//多维数组遍历
	for _, v := range a4 {
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}

	//数组是值类型，不是引用类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[2] = 2
	fmt.Println(b1, b2)

	a11 := [...]int{1, 3, 5, 7, 9}
	//计算数组中所有元素的和
	sum := 0
	for _, v := range a11 {
		sum = sum + v
	}
	fmt.Println(sum)
	//计算数组中所有元素的总两两相加的和为8的下标

	for i := 0; i < len(a11); i++ {
		for j := i + 1; j < len(a11); j++ {
			if a11[i]+a11[j] == 8 {
				fmt.Printf("%d %d \n", i, j)
			}
		}
	}

}

package main

import "fmt"

/*
	new, make 区别
	new 分配内存， make 提供初始化
	数据类型
 */

func main()  {
	fmt.Println("We are all consenting adults here")

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 2
	fmt.Println(arr1, arr2)

}

package main

import "fmt"

func main(){

	// go变量声明的方式 

	//方法1 :指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i = " ,i)


	// 方法2 : 根据值自行判断变量类型（类型推导）
	var num = 10.11
	fmt.Println("num = " , num)

	// 方法3 : 省略var ，注意 :=左边的变量不能是已经声明过的，否则会导致编译错误
	name := "lily"
	fmt.Println("name = ",name)


}
package main

import "fmt"

func main(){

	var num1 float32  = 43.000901
	fmt.Println("num1 = " ,num1)

	var num2 float64 = 43.000901
	fmt.Println("num2 = ", num2)

	//  Golang的浮点型默认声明为  float64类型。
	var num3 = 23.222
	fmt.Printf("num3 的数据类型是 %T \n",num3)


	//十进制形式
	var num4  = 5.23
	var num5 = .23  //0.23
	fmt.Println("num4 = ", num4,"num5 = ", num5)
	// 科学计数法形式
	var num6 = 5.1234e2		// //5.1234 * 10的2次方
	var num7 = 5.1234e-2  //5.1234 / 10的2次方
	fmt.Println("num6 = ", num6,"num7 = ", num7)

}
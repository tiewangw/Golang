package main

import "fmt"

func main(){

	
	var a byte = 'a'  //字符a ，用单引号
	fmt.Println("a = ",a)		// 97 ASCII码
	fmt.Printf("a =  %c \n",a)	//a   %c :a的字符值

	var b byte = '0'	 //字符0
	fmt.Println("b = ",b)		// 48	ASCII码
	fmt.Printf("b = %c \n",b)	//0

	var c = 'c'			 //字符c
	fmt.Println("c = ",c)		// 99

// ----------------------------------------------
	var a1 = "a"		//字符串a，双引号
	fmt.Println("a1 = ",a1) //a

	var b1 = 0			//数字0
	fmt.Println("b1 = ",b1)		//0

	var c1 = "c"		//字符串c
	fmt.Println("c1 = ",c1)		//c

	// var d byte = '北'
	//fmt.Println("d = ", d)    // constant 21271 overflows byte
	var d int = '北'
	fmt.Printf("北 的ASCII码是 %d ",d)



}
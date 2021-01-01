package main

import "fmt"

func main(){

	var i int  = 10
	fmt.Println("i的地址是",&i)	//i的地址是 0xc04204e058

	var prt *int = &i
	fmt.Println("prt的地址是",&prt)	//prt的地址是 0xc042004030
	fmt.Println("prt的值是",*prt)	//prt的值是 10

	// 修改prt的值同时i的值也会变
	*prt = 100
	fmt.Println("i = ",i)	//i =  100
}


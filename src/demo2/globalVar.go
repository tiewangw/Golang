package main

import "fmt"

var  var1 = 32
var	 var2 = "yryr"
var  var3 = 32.43

// 一次性声明多个全局变量
var(
	var4 = 534
	var5 = "cat"
)

func main(){

	fmt.Println("var1=",var1,"var2=",var2,"var3=",var3)
	fmt.Println("var4=",var4,"var5=",var5)
}
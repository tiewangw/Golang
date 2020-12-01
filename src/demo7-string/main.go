package main

import(
	"fmt"
)

func main(){
	
	var a = "rfere"
	fmt.Println("a = ",a)

	//使用的是反引号
	var b = `
			import(
				"fmt"
			)

			func main(){
				
				var a = "rfere"
				fmt.Println(a)

			
	
	`
	fmt.Println("b = " ,b)

	
	//注意：拼接字符串 + 要留在上一行
	var c = "hello" + "world" + 
			"hello" + "world" +
			"hello" + "world"
	fmt.Println("c = " ,c)		

}
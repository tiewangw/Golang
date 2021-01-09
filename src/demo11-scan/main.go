package main

import "fmt"

func main(){

	var name string
	var sal float32
	var age int

	fmt.Println("请输入姓名、年龄、收入，以空格隔开")
	fmt.Scanf("%s %d %f",&name,&age,&sal)
	fmt.Printf("姓名%v \n年龄%v \n收入%v",name,age,sal)

}
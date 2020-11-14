package main

import "fmt"

func main(){

	// 多变量声明 

	//方法1 :指定变量类型，声明后若不赋值，使用默认值
	var n1,n2,n3 int
	fmt.Println("n1= " ,n1 , "n2= ", n2, "n3=",n3)



	// 方法2 : （类型推导）
	var n4 ,name , n5 = 10.11,"hhh",99
	fmt.Println("n4= " ,n4 , "name= ", name, "n5=",n5)

	// 方法3 : 省略var
	n6 , name2 ,n7  :=  43 , "lily" ,7.4
	fmt.Println("n6= " ,n6 , "name2= ", name2, "n7=",n7)


}
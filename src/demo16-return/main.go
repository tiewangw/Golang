package main

import "fmt"

func test(n1 int,n2 int) (int,int)  {
	sum := n1 + n2
	sub := n1 - n2
	return sum,sub
}

func main()  {
	
	sum,sub := test(2,3)
	fmt.Println("sum = ",sum)
	fmt.Println("sub = ",sub)

	// _占位符
	res1,_ := test(12,13)
	fmt.Println("res1 = ",res1)
}
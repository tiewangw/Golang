package main

import "fmt"
// defer
func sum(n1 int,n2 int) int{

	defer fmt.Println("n1  = ",n1)
	defer fmt.Println("n2  = ",n2)

	n1++
	n2++
	n3 := n1 + n2
	fmt.Println("n3 = ",n3)

	return n3
	
}

func main(){
	
	a := sum(10,20)
	fmt.Println("a = ",a)
}
package main

import "fmt"

//闭包
func addClosure() func (int) string {

	var  str  =  "hello"

	return func (n int) string{
		str += string(n)
		return str
	}

}


func main(){

	f := addClosure()
	fmt.Println(f(36))  //hello$
	fmt.Println(f(36))	//hello$$
	fmt.Println(f(36))	//hello$$$

}
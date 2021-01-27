package main

import "fmt"


func recursion(n int) {

	if n > 2 {
		n--
		recursion(n)
	}
	fmt.Println("n = ", n)
}

func main()  {
	
	recursion(4)
	// 输出结果
	// n =  2
	// n =  2
	// n =  3
}
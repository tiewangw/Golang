package main

import "fmt"

func main()  {
	
	var a int32 = 20
 
	switch a {
		case 20,30:
			fmt.Println("a < 100")
		//	fallthrough	//switch穿透
		case 100:
			fmt.Println("a >= 100")
		default :
			fmt.Println("a 是一个int")
		
	}
}
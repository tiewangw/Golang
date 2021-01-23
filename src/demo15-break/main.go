package main

import (
	"fmt"
	"math/rand"
	"time"
)
// for   break
func main(){

	var n int
	for  {
		rand.Seed(time.Now().UnixNano())
		res := rand.Intn(100)
		fmt.Println(res)
		n++
		if res == 99 {
			break
		}
	}
	
	fmt.Printf("生成99一共用了%d次",n)
}




package main

import(
	"fmt"
	"unsafe"

)

func main(){
	
	var a = false;
	fmt.Println(a)

	// bool占用1个字节
	fmt.Println("a占用的字节= ",unsafe.Sizeof(a))


}
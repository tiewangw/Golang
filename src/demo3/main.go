
package main

import(
	"fmt"
	"unsafe"
) 

func main (){


	var a int8  = 127
	fmt.Println(a)	

	var b int16 = 128
	fmt.Println(b)


	var c  uint16 = 255
	fmt.Println(c)

	var d int16 = 432;
	//判断d的数据类型和占用的字节数
	// fmt.Printf 可以做格式化输出
	fmt.Printf("d 的类型 %T ，d 占用的字节数是 %d \n",d,unsafe.Sizeof(d))  


	// Golang程序中整型变量在使用时，遵守保小不保大的原则，
	// 即：在保证程序正确运行下，尽量使用占用空间小的数据类型。【如：年龄】
	var age byte = 90
	fmt.Println(age)

}



package main

import "fmt"

func main(){

	var num1 int32 = 456

	var num2  float32  = 43.1

	var b1 bool = false

	var b2 byte = 'f'

	var str string

	// 基本数据类型转换成string 方法1：使用fmt.Sprintf

	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="456"

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="43.099998"

	str = fmt.Sprintf("%t",b1)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="false"

	str = fmt.Sprintf("%c",b2)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="f"


}



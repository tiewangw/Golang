
package main

import(
	"fmt"
	"strconv"
)

func main(){

	var num1 int32 = 456

	var num2  float64  = 43.123

	var b1 bool = true

	var b2 byte = 'a'

	var str string

	// 基本数据类型转换成string 方法2：使用strconv包

	str = strconv.FormatInt(int64(num1),10)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="456"

	// 5表示小数位数，64 表示float64
	str = strconv.FormatFloat(num2,'f',5,64)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="43.12300"

	str = strconv.FormatBool(b1)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="true"

	str = strconv.FormatByte(b2)
	fmt.Printf("str type %T str=%q\n",str,str)	//str type string str="f"


}



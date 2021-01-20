package main

import "fmt"

func main()  {
	
	
	var str string = "hello,worldb北京"

	//方法：北京乱码
	for i := 0 ;i < len(str);i++ {
		fmt.Printf("i = %d, v = %c \n",i,str[i])
	}
	fmt.Println("--------------")
	//方法2：用range
	str2 := []rune(str)
	for j, v := range str {
		fmt.Printf("j = %d, v = %c \n",j,v)
	}
	fmt.Println("--------------")
	//方法3：用切片rune
	for k := 0 ;k < len(str2);k++ {
		fmt.Printf("k = %d, v = %c \n",k,str2[k])
	}
}
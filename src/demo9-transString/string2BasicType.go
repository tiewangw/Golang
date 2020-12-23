
package main

import (
	"strconv"
	"fmt"
)

func main() {


		var str string = "true"

		var b bool = true

		b,_ = strconv.ParseBool(str)
		fmt.Printf("b type %T , b=%v\n",b,b)	//b type bool , b=true




}

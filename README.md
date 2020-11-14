# Golang
golang官网 https://golang.google.cn/



golang中文文档  https://studygolang.com/pkgdoc



#### Golang环境变量配置及其作用:

​	GOROOT:指定  go sdk安装目录

​	Path:指令   sdk\bin目录：

​	GOPATH:就是  golang工作目录：我们的所有项目的源码都这个目录下。



#### Golang程序开发的注意事项：

​		1)  Go源文件以  "go"为扩展名。

​		2)  Go应用程序的执行入口是main()函数。这个是和其它编程语言（比如   java/c）一样。

​		3)  **Go语言严格区分大小写。**

​		4)  Go方法由一条条语句构成，**每个语句后不需要分号**  (Go语言会在每行后自动加分号)，这也体现出 Golang的简洁性。

​		5)  Go编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一个，否则会报错。

​		6)  go定义的变量，或者import包，必须使用，如果没有使用就会报错。



#### Golang规范的代码风格：

​		1)	Go官方推荐使用行注释来注释整个方法和语句。

​		2)	使用gofmt来进行格式化 。

​		3)	运算符两边习惯性各加一个空格。比如：2  + 4 * 5。



#### Golang程序的编译、运行：

编译：go build源码   = 》生成一个二进制的可执行文件

运行：方法1.对可执行文件运行   xx.exe    ./可执行文件   

​			方法2. go run源码



#### Go变量使用的3种方式

方法1 :指定变量类型，声明后若不赋值，使用默认值

方法2 : 根据值自行判断变量类型（类型推导）

方法3 : 省略var ，注意 :=左边的变量不能是已经声明过的，否则会导致编译错误

![1605189597917](C:\Users\15761\AppData\Roaming\Typora\typora-user-images\1605189597917.png)

运行结果

![1605189634131](C:\Users\15761\AppData\Roaming\Typora\typora-user-images\1605189634131.png)



##### 多变量声明

![1605190391445](C:\Users\15761\AppData\Roaming\Typora\typora-user-images\1605190391445.png)

运行结果

![1605190429099](C:\Users\15761\AppData\Roaming\Typora\typora-user-images\1605190429099.png)



##### 一次声明多个全局变量

![1605190957778](C:\Users\15761\AppData\Roaming\Typora\typora-user-images\1605190957778.png)

运行结果

![1605191002600](C:\Users\15761\AppData\Roaming\Typora\typora-user-images\1605191002600.png)




































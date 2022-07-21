// 使用整数
package main

import "fmt"

func main() {

	// 使用8位无符号整型
	var X uint8 = 225
	fmt.Println(X+1, X)

	// 使用16位有符号整型
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	a := 20.45
	b := 34.89

	//两个浮点数相减
	c := b - a

	//显示结果
	fmt.Printf("结果: %f", c)

	//显示c变量的类型
	fmt.Printf("\nc的类型是 : %T", c)

	//用于存储字符串的str变量
	str := "nhooo"

	//显示字符串的长度
	fmt.Printf("字符串的长度:%d", len(str))

	//显示字符串
	fmt.Printf("\n字符串是: %s", str)

	// 显示str变量的类型
	fmt.Printf("\nstr的类型是: %T", str)

	//变量声明和初始化
	//显式类型
	var myvariable1 = 20
	var myvariable2 = "nhooo"
	var myvariable3 = 34.80

	// Display the value and the
	// type of the variables
	fmt.Printf("myvariable1的值是 : %d\n", myvariable1)

	fmt.Printf("myvariable1的类型是 : %T\n", myvariable1)

	fmt.Printf("myvariable2的值是 : %s\n", myvariable2)

	fmt.Printf("myvariable2的类型是 : %T\n", myvariable2)

	fmt.Printf("myvariable3的值是 : %f\n", myvariable3)

	fmt.Printf("myvariable3的类型是 : %T\n", myvariable3)

	const A = "GFG"
	var B = "nhooo"

	//拼接字符串
	var helloWorld = A + " " + B
	helloWorld += "!"
	fmt.Println(helloWorld)

	//比较字符串
	fmt.Println(A == "GFG")
	fmt.Println(B < A)
}

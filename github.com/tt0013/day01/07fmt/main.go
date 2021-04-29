package main

import "fmt"

//fmt占位符
func main()  {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%#v\n", s)
}

//Go语言中字符串是用双引号包裹的！！！
//Go语言中单引号包裹的是字符！！！
//字符串
// s := "Hello"
//单独的字符，汉字，符号表示一个字符
// c1 := 'h'
// c2 := '1'
// c3 := '汉'
//字节: 1字节=8Bit(8个二进制)
//1个字符'A'=一个字节
//1个utf8编码的汉字'汉' = 一般占3个字节
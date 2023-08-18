package main

import "fmt"

//导入的包要使用，要不编译会报错
//import "os"

// 常量声明
const PI = 3.1415926
const IP = "127.0.0.1"

// 全局变量声明
var name = "aaron.qiu"

// 一般类型声明
type age int

// struct 声明
type person struct{}

// 接口声明
type iUser interface{}

func main() {
	//go语言通过函数首字母的大小写来区分函数是public ,还是private
	fmt.Printf("hello, world\n")
	var str1 string = "small"
	fmt.Println(str1)
	str2 := "solo" //自动解析出类型,只能在函数内部定义
	fmt.Println(str2)

	a := 1
	if a == 1 {
		fmt.Println("if a equal one is true")
	} else if a == 2 {
		fmt.Println("if a equal two is true")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}

	fmt.Printf("hello, world\n") //
}

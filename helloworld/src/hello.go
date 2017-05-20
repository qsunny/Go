package main

import "fmt"

//导入的包要使用，要不编译会报错
//import "os"

//常量声明
const PI = 3.1415926

//全局变量声明
var name = "aaron.qiu"

//一般类型声明
type age int

//struct 声明
type person struct{}

//接口声明
type iUser interface{}

func main() {
	fmt.Printf("hello, world\n") //
}

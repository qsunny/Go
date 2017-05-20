package main

//import way 1
import "fmt"

//import way 2
import (
	"math"
	"math/cmplx"
	"math/rand"
)

//var关键字变量声明,可以在package,func
//var c, python, java bool
var c, python, java = true, false, "no!"

//Go's basic types are
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//函数定义
func add(x int, y int) int {
	return x + y
}

//多个返回值函数
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var n int
	fmt.Println(n, c, python, java)
	/*
	 *在函数体内通过:=隐含声明变量，在函数体外不能使用这种方式
	 *必须通过var,func关键字
	 */
	k := 3

	fmt.Println(k)
	for i := 0; i < 3; i++ {
		fmt.Println("My Favorite number is", rand.Intn(10))
	}
	fmt.Println("Now you have %g problems.", math.Sqrt(7))

	fmt.Println(math.Pi) //math.pi 小写是私有变量，没有暴露不能访问
	//函数调用
	fmt.Println(add(12, 16))

	a, b := swap("nala", "pikaqiu")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//类型转换
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
	fmt.Printf("v is of type %T\n", i)
	//或者下面这种方式
	ii := 42
	ff := float64(i)
	uu := uint(f)
	fmt.Println(ii, ff, uu)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

/*
https://mp.weixin.qq.com/s/k5CpO6vVSMP4tOYAcAXcqg
 */
func main() {
	fmt.Println("Hello,欢迎你GoLang")
	beyondHello()
	learnTypes()
	learnFlowControl() // 回到流程控制

	learnFunctionFactory() // 返回函数的函数多棒啊
	learnDefer()      // 对defer关键字的简单介绍
	learnInterfaces() // 好东西来了！

	learnVariadicParams("great", "learning", "here!")
	learnErrorHandling()

	learnConcurrency()

	learnWebProgramming() // Go很适合web编程
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)

}

func learnMultiple(x, y int) (sum, prod int) {
	return x+y, x*y //返回两个值
}

// 内置变量类型和关键词
func learnTypes()  {
	str := "少说话、多干活!"

	s2 := `这是一个
可以换行的字符串`

	fmt.Println(str,s2)

	// 非ascii字符。Go使用UTF-8编码。
	g := 'Σ' // rune类型，int32的别名，使用UTF-8编码

	f := 3.14195 // float64类型，IEEE-754 64位浮点数
	c := 3 + 4i  // complex128类型，内部使用两个float64表示

	fmt.Println(g,f,c)

	var u uint = 7
	var pi float32 = 22. /7

	fmt.Println(u,pi)

	// 字符转换
	n := byte('\n') // byte是uint8的别名

	fmt.Println(n)

	var a4 [4] int
	a3 := [...]int{3,1,5,4}

	fmt.Println(a4,a3)

	s3 := []int{4,5,7}
	s4 := make([]int,4)
	var d2 [][]float64
	bs := []byte("a slice")

	fmt.Println(s3,s4,d2,bs)

	s := []int{1,2,3}
	s = append(s,4,5,6)

	fmt.Println(s)

	// 除了向append()提供一组原子元素（写死在代码里的）以外，我们
	// 还可以用如下方法传递一个slice常量或变量，并在后面加上省略号，
	// 用以表示我们将引用一个slice、解包其中的元素并将其添加到s数组末尾。
	s = append(s,[]int{7,8,9}...)

	fmt.Println(s)

	p, q := learnMemory()       // 声明p,q为int型变量的指针
	fmt.Println(*p, *q)         // * 取值

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	fmt.Println(m)

	// 在Go语言中未使用的变量在编译的时候会报错，而不是warning。
	// 下划线 _ 可以使你“使用”一个变量，但是丢弃它的值。
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs
	// 通常的用法是，在调用拥有多个返回值的函数时，
	// 用下划线抛弃其中的一个参数。下面的例子就是一个脏套路，
	// 调用os.Create并用下划线变量扔掉它的错误代码。
	// 因为我们觉得这个文件一定会成功创建。
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "这句代码还示范了如何写入文件呢")
	file.Close()

	// 输出变量
	fmt.Println(s, c, a4, s3, d2, m)

}

// 和其他编程语言不同的是，go支持有名称的变量返回值。
// 声明返回值时带上一个名字允许我们在函数内的不同位置
// 只用写return一个词就能将函数内指定名称的变量返回
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // z is implicit here, because we named it earlier.
}

// Go全面支持垃圾回收。Go有指针，但是不支持指针运算。
// 你会因为空指针而犯错，但是不会因为增加指针而犯错。
func learnMemory() (p, q *int) {
	// 返回int型变量指针p和q
	p = new(int)    // 内置函数new分配内存
	// 自动将分配的int赋值0，p不再是空的了。
	s := make([]int, 20)    // 给20个int变量分配一块内存
	s[3] = 7                // 赋值
	r := -2                 // 声明另一个局部变量
	return &s[3], &r        // & 取地址
}

func expensiveComputation() int {
	return 1e6
}

func learnFlowControl() {
	// If需要花括号，括号就免了
	if true {
		fmt.Println("这句话肯定被执行")
	}
	// 用go fmt 命令可以帮你格式化代码，所以不用怕被人吐槽代码风格了，
	// 也不用容忍别人的代码风格。
	if false {
		// pout
	} else {
		// gloat
	}
	// 如果太多嵌套的if语句，推荐使用switch
	x := 1
	switch x {
	case 0:
	case 1:
		// 隐式调用break语句，匹配上一个即停止
	case 2:
		// 不会运行
	}

	// 和if一样，for也不用括号
	for x := 0; x < 3; x++ { // ++ 自增
		fmt.Println("遍历", x)
	}
	// x在这里还是1。为什么？

	// for 是go里唯一的循环关键字，不过它有很多变种
	for { // 死循环
		break    // 骗你的
		continue // 不会运行的
	}

	// 用range可以枚举 array、slice、string、map、channel等不同类型
	// 对于channel，range返回一个值，
	// array、slice、string、map等其他类型返回一对儿
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// 打印map中的每一个键值对
		fmt.Printf("索引：%s, 值为：%d\n", key, value)
	}
	// 如果你只想要值，那就用前面讲的下划线扔掉没用的
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("你是。。 %s\n", name)
	}

	// 和for一样，if中的:=先给y赋值，然后再和x作比较。
	if y := expensiveComputation(); y > x {
		x = y
	}

	// 闭包函数
	xBig := func() bool {
		return x > 100 // x是上面声明的变量引用
	}
	fmt.Println("xBig:", xBig()) // true （上面把y赋给x了）
	x /= 1e5                     // x变成10
	fmt.Println("xBig:", xBig()) // 现在是false

	// 除此之外，函数体可以在其他函数中定义并调用，
	// 满足下列条件时，也可以作为参数传递给其他函数：
	//   a) 定义的函数被立即调用
	//   b) 函数返回值符合调用者对类型的要求
	fmt.Println("两数相加乘二: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Called with args 10 and 2
	// => Add + double two numbers: 24

	// 当你需要goto的时候，你会爱死它的！
	goto love
love:


}

func learnFunctionFactory() {
	// 空行分割的两个写法是相同的，不过第二个写法比较实用
	fmt.Println(sentenceFactory("原谅")("当然选择", "她！"))

	d := sentenceFactory("原谅")
	fmt.Println(d("当然选择", "她！"))
	fmt.Println(d("你怎么可以", "她？"))
}

func sentenceFactory(myString string) func(before,after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, myString, after)
	}
}

func learnDefer() (ok bool) {
	// defer表达式在函数返回的前一刻执行
	defer fmt.Println("defer表达式执行顺序为后进先出（LIFO）")
	defer fmt.Println("\n这句话比上句话先输出，因为")
	// 关于defer的用法，例如用defer关闭一个文件，
	// 就可以让关闭操作与打开操作的代码更近一些
	return true
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

// 定义pair类型的方法，实现Stringer接口。
func (p pair) String() string { // p被叫做“接收器”
	// Sprintf是fmt包中的另一个公有函数。
	// 用 . 调用p中的元素。
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}




func learnInterfaces() {
	// 花括号用来定义结构体变量，:=在这里将一个结构体变量赋值给p。
	p := pair{3,4}
	fmt.Println(p.String()) // 调用pair类型p的String方法

	var i Stringer // 声明i为Stringer接口类型
	i = p // 有效！因为p实现了Stringer接口（类似java中的塑型）
	// 调用i的String方法，输出和上面一样
	fmt.Println(i.String())

	// fmt包中的Println函数向对象要它们的string输出，实现了String方法就可以这样使用了。
	// （类似java中的序列化）
	fmt.Println(p) // 输出和上面一样，自动调用String函数。
	fmt.Println(i) // 输出和上面一样。

}

func learnVariadicParams(myStrings ...interface{}) {
	// 枚举变长参数列表的每个参数值
	// 下划线在这里用来抛弃枚举时返回的数组索引值
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// 将可变参数列表作为其他函数的参数列表
	fmt.Println("params:", fmt.Sprintln(myStrings...))


}

func learnErrorHandling() {
	// ", ok"用来判断有没有正常工作
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok { // ok 为false，因为m中没有1
		fmt.Println("别找了真没有")
	} else {
		fmt.Print(x) // 如果x在map中的话，x就是那个值喽。
	}
	// 错误可不只是ok，它还可以给出关于问题的更多细节。
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// 输出"strconv.ParseInt: parsing "non-int": invalid syntax"
		fmt.Println(err)
	}
}

// c是channel类型，一个并发安全的通信对象。
func inc(i int, c chan int) {
	c <- i + 1 // <-把右边的发送到左边的channel。
}

// 我们将用inc函数来并发地增加一些数字。
func learnConcurrency() {
	// 用make来声明一个slice，make会分配和初始化slice，map和channel。
	c := make(chan int)

	// 用go关键字开始三个并发的goroutine，如果机器支持的话，还可能是并行执行。
	// 三个都被发送到同一个channel。
	go inc(0, c) // go is a statement that starts a new goroutine.
	go inc(10, c)
	go inc(-805, c)
	// 从channel中读取结果并打印。
	// 打印出什么东西是不可预知的。
	fmt.Println(<-c, <-c, <-c) // channel在右边的时候，<-是读操作。

	cs := make(chan string) // 操作string的channel
	cc := make(chan chan string) // 操作channel的channel
	go func() { c <- 84 }()       // 开始一个goroutine来发送一个新的数字
	go func() { cs <- "wordy" }() // 发送给cs
	// Select类似于switch，但是每个case包括一个channel操作。
	// 它随机选择一个准备好通讯的case。
	select {
	case i := <-c: // 从channel接收的值可以赋给其他变量
		fmt.Println("这是……", i)
	case <-cs: // 或者直接丢弃
		fmt.Println("这是个字符串！")
	case <-cc: // 空的，还没作好通讯的准备
		fmt.Println("别瞎想")
	}
	// 上面c或者cs的值被取到，其中一个goroutine结束，另外一个一直阻塞。
}

func learnWebProgramming() {
	// ListenAndServe第一个参数指定了监听端口，第二个参数是一个接口，特定是http.Handler。
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // 不要无视错误。
	}()

	requestServer()
}

// 使pair实现http.Handler接口的ServeHTTP方法。
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 使用http.ResponseWriter返回数据
	w.Write([]byte("Y分钟golang速成!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\n服务器消息： `%s`", string(body))
}
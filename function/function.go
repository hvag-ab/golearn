package function

/**
 *  datetime: 2017-05-23
 *  Go语言 - 函数
 *  0. 函数的零值是nil
 *  1. Go中函数声明包含：函数名称，参数列表, 参数名称，返回值名称，返回值类型。
 *  2. Go函数可以有多个返回结果。在函数有多个返回结果情况下，要么全部省略返回值名称，要么全部加上返回值名称
 *  3. 如果函数结果有名称，那么函数在被调用时，他们的变量就会被隐显声明。
 *  4. 在Go中习惯性的将错误类型作为函数结果列表中的最后一员。errors.New("异常")
 *  5. 函数在Go中是一等类型，函数可以当做其他函数的参数，也可以当做结果。
 *  6. Go中的函数闭包：内层函数引用了外层函数环境的变量，其返回值也是一个函数。
	7. go中函数传值是值拷贝 所以不会改变原来值  一般传递指针 指针改变原来数值
	
	//在函数外部定义的变量成为全局变量
//全局变量和局部变量可以重名
//全局变量可以在项目中所有文件进行使用
//在任意位置修改全局变量的值都会影响其他位置使用
//全局变量存储在数据区

var a int =100
//全局变量定义时不能使用自动推到类型
//c:=100//err

 */

 import (
	 "errors"
	 "fmt"
	 "log"
 )

 func init(){//在main函数之前执行初始化函数  用来配置一些东西 比如日志
    log.SetPrefix("【Hvag】")
    log.SetFlags(log.LstdFlags | log.Lshortfile )
}
 
 // 别名
 type myint int
 
 /**
  *   声明一个函数
  *   函数名：f1,  参数名称：i, 参数类型：myint,  返回结果类型int
  */
 func F1(i myint) myint {
	 i++
	 r := i + 1
	 return r
 }
 
 
 /**
  *   1. 函数中有异常信息一般作为函数结果列表中最后一个参数抛出。
  */
 func Divide(i, j int) (result int, err error) {
 
	 if j == 0 {
		 err = errors.New("divison by zero")
		 return //等价于返回 0（int默认值），err
	 }
	 result = i / j
	 return //等价于返回 result，nil 
 }
 
 
 //可变参数，最多只能有一个可变参数，可变参数只能放到函数参数的末尾
 // Sum(1,2,3,4,5)
 func Sum(nums ...int) int {
	 total := 0
	 for _, num := range nums {
		 total += num
	 }
	 return total
 }
 

 //函数作为参数
 func Callback(y int, f func(int, int)) {
	 f(y, 2) // this becomes Add(1, 2)
 }
 // Add(a, b int)
 //调用函数 Callback(2, Add(2，3))
 
// 指针作为参数
func F4(a *int, b *int) {//可以看作 var a *int = &a 所以需要传递&a 作为参数
	//temp := *a
	//*a = *b
	//*b = temp
	*a, *b = *b, *a
	fmt.Println("a",*a,a)//因为传递过来的是一个地址 所以需要取地址也就是指针 需要加*取地址对应的值
}
// var a,b int = 2,2
// F4(&a, &b)

// 递归
func Fib(n int) int{

	if n == 1 || n==2{
		return 1
	}else{
		return Fib(n-1)+Fib(n-2)
	}
}

// 匿名函数
var Fn = func(i int) int{
	return i+2
}

// Fn(3)

//定义一个函数squre , 返回值为一个匿名函数, 也就是通常说的闭包
func Square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
// var i int = Square()()

//defer 在函数返回值后在调用 调用完后 在return 所以 先返回x=10 在调用defer 最后在执行returnp0
//defer 用在关闭资源 比如close 文件 数据库 
func Defer() int {

	x := 10
	defer func(x int) {
		x++
	}(2)
	return x
}

// 装饰器模式
type Decoer func(i int, s string) bool

func foo(i int, s string) bool {
	fmt.Printf("=== foo ===\n")
	return true
}

func withTx(fn Decoer) Decoer {
	return func(i int, s string) bool {
		fmt.Printf("=== start tx ===\n")
		result := fn(i, s)
		fmt.Printf("=== commit tx ===\n")

		return result
	}
}

// foo := withTx(foo)
// foo(1, "hello")

// 任意接口
func Any(v interface{})  { //空接口用来传递任意类型的参数

    if v2, ok := v.(string);ok{
        println(v2)
    }else if v3,ok2:=v.(int);ok2{
        println(v3)
    }
}


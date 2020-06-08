/**
 *  datetime: 2017-05-29
 *  Go语言 - 匿名函数（闭包）
 *  函数返回的匿名函数属于引用类型
 *  内部函数访问或者改变外部函数的变量
 */
package main

import "fmt"

//定义一个函数squre , 返回值为一个匿名函数, 也就是通常说的闭包

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := square()
	fmt.Println(f()) //将x的状态记录到了下来
	fmt.Println(f()) // 使用上一次执行得到的x值计算。
	fmt.Println(f())
	hvag()
}

func hvag() {
    for i := 0; i < 4; i++ {
        g := func(i int) { fmt.Printf("%d ", i) }//匿名函数
        g(i)
        fmt.Printf(" - g is of type %T and has value %v\n", g, g)
    }
}
//$ go run anonymous_func.go

//1
//4
//9

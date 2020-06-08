package main

import "fmt"

func test3(a int, b int) {
	a, b = b, a
}
func test4(a *int, b *int) {//可以看作 var a *int = &a 所以需要传递&a 作为参数

	//temp := *a
	//*a = *b
	//*b = temp
	*a, *b = *b, *a
	fmt.Println("a",*a,a)//因为传递过来的是一个地址 所以需要取地址也就是指针 需要加*取地址对应的值
}
func main() {
	a := 10
	b := 20
	//变量作为函数参数是值传递
	//test3(a, b)
	//指针作为函数参数是地址传递
	test4(&a, &b)
	fmt.Println(a, b)
}

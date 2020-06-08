package main

import "fmt"

func main() {

	var a int = 10

	fmt.Printf("%p\n", &a)

	//将a的地址赋值给一个指针变量
	//指针* 表示一级指针
	var p *int = &a //定义一个p变量 指向a int变量的内存地址
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n",&p)
	fmt.Printf("Here is the string *p: %d\n", *p) // prints string *p就取 a的值

	//通过指针间接修改变量的值
	*p = 100

	fmt.Println(a)

	fmt.Printf("%T\n", p)
	//fmt.Printf("%T\n",&p)

	var hv *int
	hv = new(int)
	*hv = 1
	fmt.Printf("%T\n",&hv,*hv)
}

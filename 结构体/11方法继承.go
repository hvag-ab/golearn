package main

import "fmt"

type person struct {
	id   int
	name string
	sex  string
	age  int
}

type student struct {
	person

	score int
	addr  string
}

//接收者的类型如果不相同 就是不同的方法
//函数名 可以和方法重名
//func Print() {
//	fmt.Print("hello")
//}

func (p *person) Print() {
	fmt.Printf("大家好，我是%s,我今年%d岁\n", p.name, p.age)
}

func main() {

	var stu student=student{person{101,"lh","男",10},100,"bjj"}

	//子类可以从父类继承结构体成员  也可以继承父类的方法
	//stu.Print()
	stu.person.Print()


	var per person=person{102,"挣钱","男",10}

	per.Print()

}

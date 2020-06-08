package main

import "fmt"

type person2 struct {
	name string
	age  int
	sex  string
}

type student2 struct {
	//指针匿名字段
	*person2

	score int
	id    int
	addr  string

	string  //匿名字段
}

func main() {



	var stu student2

	//stu.person2指针变量位空 nil
	//通过new创建空间 存储时数据
	stu.person2 = new(person2)

	stu.name = "lh"
	stu.sex = "男"
	stu.age = 20

	stu.id = 102
	stu.score = 100
	stu.addr = "bb"

	stu.string = "Q"

	fmt.Println(stu)
	fmt.Println(stu.name)
	fmt.Println(stu.sex)
	fmt.Println(stu.age)
	fmt.Println(stu.string)
}

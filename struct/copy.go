package struct_

import "fmt"

/*
深拷贝 浅拷贝
二、本质区别：
是否真正获取（复制）对象实体，而不是引用。

三、如何理解？
这里举个例子，比如P2复制了P1，修改P1属性的时候，观察P2的属性是否会产生变化

1、P2的属性变化了，说明这是浅拷贝，堆中内存还是同一个值。

p2=&p1 // 浅拷贝，p2为指针，p1和p2共用一个内存地址
2、P2的属性没变化，说明这是深拷贝，堆中内存是不同的值了。

p2=p1 // 深拷贝，生成两个内存地址
*/


// 定义一个Robot结构体
type Robot struct {
	Name  string
	Color string
	Model string
 }

 func main45() {
	fmt.Println("深拷贝 内容一样，改变其中一个对象的值时，另一个不会变化。")
	robot1 := Robot{
	   Name:  "小白-X型-V1.0",
	   Color: "白色",
	   Model: "小型",
	}
	robot2 := robot1 
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, &robot2)
 
	fmt.Println("修改Robot1的Name属性值")
	robot1.Name = "小白-X型-V1.1"
 
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, &robot2)

	fmt.Println("浅拷贝 内容和内存地址一样，改变其中一个对象的值时，另一个同时变化。")
	robot3 := &robot1
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 3：%s\t内存地址：%p \n", robot3, robot3)

	fmt.Println("在这里面修改Robot1的Name和Color属性")
	robot1.Name = "小黑-X型-V1.2"
	robot1.Color = "黑色"

	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 3：%s\t内存地址：%p \n", robot3, robot3)

	fmt.Println("浅拷贝 使用new方式")
	robot8 := new(Robot)
	robot8.Name = "小白-X型-V1.0"
	robot8.Color = "白色"
	robot8.Model = "小型"
 
	robot9 := robot8
	fmt.Printf("Robot 8：%s\t内存地址：%p \n", robot8, robot8)
	fmt.Printf("Robot 9：%s\t内存地址：%p \n", robot9, robot9)
 
	fmt.Println("在这里面修改Robot8的Name和Color属性")
	robot8.Name = "小蓝-X型-V1.3"
	robot8.Color = "蓝色"
 
	fmt.Printf("Robot 8：%s\t内存地址：%p \n", robot8, robot8)
	fmt.Printf("Robot 9：%s\t内存地址：%p \n", robot9, robot9)
 
 }
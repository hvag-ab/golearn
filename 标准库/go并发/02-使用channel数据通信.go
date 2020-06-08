package main

import (
	"fmt"
	"time"
)


func main()  {

	// 在主、子go程之间进行数据传递。 指定了 读写数据的先后顺序。 子先，主后
	ch := make(chan int)

	// 在 主子 go 程使用 stdout，指定先后顺序， 子先，主后
	ch2 := make(chan string)

	// 创建匿名子go程
	go func() {
		fmt.Println("子 go 程 start ....")
		for i:=0; i<5;i++ {
			ch <- i
			fmt.Println("子go程 i = ", i)
			ch2 <- "write finish"
		}
	}()
	time.Sleep(time.Second * 1)		// 保证子go程 ，先执行

	for i:=0; i<5; i++ {
		num := <-ch
		<- ch2				// 读到 子go程 发送的 控制 使用stdout 的字符串，之后，在 使用stdout
		fmt.Println("主go程读到：", num)
	}
}

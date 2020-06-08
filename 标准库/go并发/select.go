package main

import (
    "fmt"
    "time"
)

// //select基本用法
// select {
// case <- chan1:
// // 如果chan1成功读到数据，则进行该case处理语句
// case chan2 <- 1:
// // 如果成功向chan2写入数据，则进行该case处理语句
// default:
// // 如果上面都没有成功，则进入default处理流程
// 如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
// 如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
// 如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

func service1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from service1"
}
func service2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from service2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go service1(ch1)
	go service2(ch2)
	
	select {       // 会发送阻塞
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}
}

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	go service1(ch1)
// 	go service2(ch2)

// 	time.Sleep(time.Second)   // 延时 1s,等待 ch1 ch2 准备就绪
	
// 	select {
// 	case s1 := <-ch1:
// 		fmt.Println(s1)
// 	case s2 := <-ch2:
// 		fmt.Println(s2)
// 	default:
// 		fmt.Println("no case ok")
// 	}
// }


//超时设定
// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	go service1(ch1)
// 	go service2(ch2)

// 	select {       // 会发送阻塞
// 	case s1 := <-ch1:
// 		fmt.Println(s1)
// 	case s2 := <-ch2:
// 		fmt.Println(s2)
// 	case <-time.After(2*time.Second):     // 等待 2s
// 		fmt.Println("no case ok")
// 	}
// }






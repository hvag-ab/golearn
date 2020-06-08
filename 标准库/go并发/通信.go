package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go sendData(ch)
    go getData(ch)  

    time.Sleep(1e9)
}

func sendData(ch chan string) {
    ch <- "Washington"
    ch <- "Tripoli"
    ch <- "London"
    ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
    // var input string
    // time.Sleep(1e9)
	for {
        input, open := <-ch
        if !open {
            break
        }
        fmt.Printf("%s ", input)
    }
}

/*
var ch1 chan string
ch1 = make(chan string)
当然可以更短： ch1 := make(chan string)。

这里我们构建一个int通道的通道： chanOfChans := make(chan int)。
通信操作符 <-
这个操作符直观的标示了数据的传输：信息按照箭头的方向流动。

流向通道（发送）

ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）

从通道流出（接收），三种方式：

int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。
*/
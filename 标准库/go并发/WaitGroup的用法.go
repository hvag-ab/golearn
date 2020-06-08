package main

import (
    "fmt"
    "sync"
)

// WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
// Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。


func main() {
    waitGroup := &sync.WaitGroup{}
    DoSomething(waitGroup)
    waitGroup.Wait() // 这里会阻塞main，直到所有的任务都完成
    fmt.Println("end")
}

func DoSomething(waitGroup *sync.WaitGroup) {
    for i:=0;i <10;i++ {
        waitGroup.Add(1)//开一个协程就添加1个
        go func(waitGroup *sync.WaitGroup) {
            fmt.Print("1-")
            defer waitGroup.Done()//协程完成后就释放一个
        }(waitGroup)
    }
}

// func main() {
//     wg := sync.WaitGroup{}
//     wg.Add(100)
//     for i := 0; i < 100; i++ {
//         go f(i, &wg)
//     }
//     wg.Wait()
// }

// // 一定要通过指针传值，不然进程会进入死锁状态  WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址：
// func f(i int, wg *sync.WaitGroup) { 
//     fmt.Println(i)
//     wg.Done()
// }
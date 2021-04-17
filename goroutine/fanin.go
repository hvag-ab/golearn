package goroutine

import (
	"fmt"
	"sync"
)

/*
我们这次试用FAN-OUT和FAN-IN，解决《Golang并发模型：轻松入门流水线模型》中提到的问题：计算一个整数切片中元素的平方值并把它打印出来。
•producer()保持不变，负责生产数据。
•squre()也不变，负责计算平方值。
•修改main()，启动3个square，这3个squre从producer生成的通道读数据，这是FAN-OUT。
•增加merge()，入参是3个square各自写数据的通道，给这3个通道分别启动1个协程，把数据写入到自己创建的通道，并返回该通道，这是FAN-IN。

*/

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range nums {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
    // FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func FANIN() {
	in := producer(1, 2, 3, 4)

    // FAN-OUT 开3个并发协程
  c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	// consumer
	for ret := range merge(c1, c2, c3) {
		fmt.Printf("%3d ", ret)
	}
	fmt.Println()
}

/*
优化FAN模式

既然FAN模式不一定能提高性能，如何优化？

不同的场景优化不同，要依具体的情况，解决程序的瓶颈。

我们当前程序的瓶颈在FAN-IN，squre函数很快就完成，merge函数它把3个数据写入到1个通道的时候出现了瓶颈，适当使用带缓冲通道可以提高程序性能，再修改下代码

•merge()中的out修改为：

out := make(chan int, 100)


*/

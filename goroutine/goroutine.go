package goroutine

/*
goroutine是go语言的并发体。在go语言里面能使用go关键字来实现并发。
go func()

1.1 概念介绍
并发

一个计算机（一个cpu）上能同时执行多项任务，让cpu在某段时间内进行分片，在某段很短时间内执行程序a，然后又迅速得切换到程序b去执行，让人们看起来就像是两个程序在同时进行，这就是并发。

进程

但是人们随之发现，cpu在切换程序的时候，如果不保存上一个程序的状态（也就是我们常说的context--上下文），直接切换下一个程序，就会丢失上一个程序的一系列状态，
于是引入了进程这个概念，用以划分好程序运行时所需要的资源。因此进程就是一个程序运行时候的所需要的基本资源单位（也可以说是程序运行的一个实体）。

并行

如果一个电脑有多个cpu，每个cpu都有进程在运行，这就是并行。

用户态与内核态

为了防止用户程序做出一些危险的指令，如关机，更改系统变量，修改别的进程数据，系统分为两种运行状态，用户态以及内核态，
用户态是我们的程序所在的状态，不能随便对内核的底层进行操作。如果我们需要使用内核的底层操作的时候，内核提供了一种调用内核的接口，我们调用这些接口也就是系统调用，在进行系统调用的时候，cpu会切换到内核态，才能执行内核的函数。

线程

人们又发现一个问题，cpu切换多个进程的时候，会花费不少的时间，因为切换进程需要切换到内核态，而每次调度需要内核态都需要读取用户态的数据，进程一旦多起来，cpu调度会消耗一大堆资源，因此引入了线程的概念，线程本身几乎不占有资源，他们共享进程里的资源，内核调度起来不会那么像进程切换那么耗费资源。

协程

但是线程还是需要内核去进行调度，切换起来也是需要把用户态的数据写入到内核态，也是需要耗费一定的计算机资源，那可以不可以将切换的调度改成我们自己控制的呢，答案是有的，协程就是把自己的调度算法交给程序（用户态）去进行管理，能以更小的资源去进行并发。

goruntine

goroutine就是一个协程例子，可以根据自身调度器进行调度，当某个gooutine调用了time.sleep方法或者channel，mutex阻塞时候，调度器会使其入睡，唤醒另一个goroutine，根本不需要进入到内核态。
gorutine有1个特性当main函数结束后，由main函数启动的gorutine也会全部消失。
Go语言中os线程和goroutine的关系

cpu执行的最小单位是os线程，所有的gorutine最终都需要被runtime调度、映射到真正的os线程上，才能被CPU执行。

1个os线程对应用户态N个goroutine。

1个go程序可以同时使用N个os线程。

goroutine和os线程是多对多的映射关系，即m:n。

gorutine的调用模型（GMP）
GPM是Go语言运行时（runtime）层面的实现的，其目的是把M个gorutine映射成N个os线程，然后再被os调度到cpu执行。

G（goroutine）里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
P（processor） 管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是1：1映射的关系， 一个groutine最终是要放到M上执行的；
P与M一般也是1：1对应的。

他们关系是：

P管理着一组G挂载在M上运行。

Processor里其中1个Gorutine长久阻塞在一个Machine上时，runtime会新建一个Machine，阻塞Gorutine所在的Processor会把剩余的Gorutine挂载在新建的Mechine上。

当被阻塞的Gorutie阻塞完成或者认为其已经死掉时 回收旧的Machine。

推理得出：1Processor管理N个gorutines-------->1个processor又被挂载1个mechine运行----------》1个mechine映射到1个os线程--------->被cpu执行掉。

 

P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能
*/
import (
  "fmt"
  "time"
  "sync"
  "runtime"
)


// 并发例子
func Base() {
    go func(){
        fmt.Println("hello")
    }()

    go func(){
       fmt.Println("world")
    }()
    time.Sleep(1 * time.Second)
    //main函数结束之后由main函数启动的gorutine也全部结束
}

// 安全退出
/*
goroutine只有在自身所在函数运行完毕，或者主函数运行完毕才会打断，因而上面的例子需要等待一秒，不然未执行完的goroutine会直接被打断。 如果我们并发的线程数量多了之后，我们不可能在main里面设置一个精确睡眠时间来评估所有的goroutine已经运行完毕然后退出。
 这时候我们可以使用sync.WaitGroup来等待所有运行的goroutine运行结束后，再来退出main函数，
 主要原理是维护一个goroutine数量的计数器，每运行一个goroutine，计数器会加+1，运行结束后，计数器会-1，
 然后调用wait方法会一直阻塞，知道计数器为0，也就是当前运行的goroutine数量为0，实例如下：
*/

func SafeWait() {
    var wg sync.WaitGroup
    // 开20个协程
    for i := 0; i < 20; i++ {
        wg.Add(1)
        //  一定要通过指针传值，不然进程会进入死锁状态  WaitGroup对象不是一个引用类型 函数是传值 相当于复制一个副本
        go func(i int, wg *sync.WaitGroup) {
            // 运行完后计数器减1 一定要在单个协程中操作
            defer wg.Done()
            time.Sleep(1 * time.Second)
            fmt.Printf("goroutine %d is running\n", i)
        }(i, &wg)
    }
    // 等待计数器为0
    wg.Wait()
}


 
func f1(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Printf("f1:%d\n", i)
    }
}
 
func f2(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        fmt.Printf("f2:%d\n", i)
    }
}
 
func main89() {
    var wg sync.WaitGroup
    //默认go的runtime会把gorutine调度到机器所有cpu上
    //不过我们可设置gorutine可以被调度到几个CUP上
    runtime.GOMAXPROCS(8)
    wg.Add(2)
    go f1(&wg)
    go f2(&wg)
    wg.Wait()
}

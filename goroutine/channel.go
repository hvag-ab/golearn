package goroutine

import (
  "fmt"
  "time"
  "sync"
  "os"
  "runtime"
)

/*
channel

channel是goroutine之间的通信机制，它可以让一个goroutine通过它给另一个goroutine发送数据，
每个channel在创建的时候必须指定一个类型，指定的类型是任意的。

创建和声明：
//声明
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan string// 声明一个传递字符串的通道
var ch3 chan []int // 声明一个传递int切片的通道
var ch4 chan map[int]int // 声明一个传递map的通道

// 下面一般用于控制通道的权限才用
// 大多用于函数的入参来作为声明
var ch1 <-chan int   // 声明一个只读整型的通道
var ch1 chan<- int   // 声明一个只写整型的通道


//创建

ch5 := make(chan int)
ch6 := make(chan int，10) // 10为channel的容量大小，当channel里面有10个元素时候，这时往里面塞会阻塞
ch7 := make(chan string)
ch8 := make(chan []int)
ch9 := make(chan map[int]int )

2、channel的一些基本操作：发送（send）、接收(receive）和关闭（close）三种操作
// 发送和接收都使用 <- 符号。

// 初始化channel
ch := make(chan int)

// 把1发送到ch中
ch <- 1

// 从ch中接收值并赋值给变量x
x := <- ch

// 优雅的接受 ok为false表示通道已经关闭
x,ok := <- ch

// 从ch中接收值，忽略结果
<-ch

// 关闭channel
close(ch)


channel的capacity是固定了，length是动态的。

只要在初始化阶段固定了channel的长度（N）。

生产者只能投放N个数据，不能让channel超载数据。

同理channel中的（N）个数据读取（消费）完了，消费者也无法再消费。


*/


// 如向channel发送数据的时候，该goroutine会一直阻塞直到另一个goroutine接受该channel的数据，
// 反之亦然，goroutine接受channel的数据的时候也会一直阻塞直到另一个goroutine向该channel发送数据
// go程开启之前使用通道
func DeadLock1() {
    ch := make(chan string)
    // 在此处阻塞，然后程序会弹出死锁的报错 因为没有接收这个通道
    ch <- "hello"
    fmt.Println("channel has send data")
}

// 一个通道在一个主go程里同时进行读和写
func DeadLock2() {
    ch := make(chan string)
    // 在此处阻塞，然后程序会弹出死锁的报错 因为没有接收这个通道 主程序是逐步运行 所以运行到这里找不到接收的 虽然下面有接收的
    // ch := make(chan string,1) 可以消除阻塞 
    ch <- "hello"
    <- ch
    fmt.Println("channel has send data")
}

func DeadLock3() {
	q := make(chan int, 2)
	<-q
}

func DeadLock4() {
	q := make(chan int, 2)
	q <- 1
	q <- 2
	q <- 3 //最多只能缓存2个 超出 死锁
}

//向已经关闭的channel中写入数据 panic
func DeadLock5() {
	q := make(chan int, 2)
	close(q)
	q <- 1
}

// for range循环的通道 必须在写数据完后要close 否则一直循环
func DeadLock6() {
	pipline := make(chan string)
	go func() {
		pipline <- "hello world"
		pipline <- "hello China"
		// close(pipline)
	}()
	for data := range pipline{
		fmt.Println(data)
	}
}

// 通道1中调用了通道2，通道2中调用通道1
func DeadLock7() {
	c1,c2:=make(chan int),make(chan int)
	go func() {
		for  {
			select{
				case <-c1:
					c2<-10
			}
		}
	}()
	for  {
		select{
		case <-c2:
			c1<-10
		}
	}
}

// 子协程 阻塞 但是主程序未堵塞 程序不会报错 主运行完后 程序就结束 不管子程序是否阻塞
func ChildDeadLock(){
	c := make(chan int)

	go func(c chan int){
		fmt.Println("into goroutine")
		c <-5
		fmt.Println("end  goroutine")

	}(c)

	time.Sleep(1 * time.Second)
	fmt.Println("exit")

}

// 使用了未初始化的channel


func DeadLock8() {
  var ch chan int // 未初始化，值为 nil  需要make
	go func(i int) {
		ch <- i
	}(1)

    fmt.Println("Result: ", <-ch)
    time.Sleep(2 * time.Second)
}

//
func DeadLock9() {

        ch1 := make(chan int)

        ch2 := make(chan int)

        go func() {

            ch1 <- 1

            ch2 <- 2

        }()

        fmt.Println(<-ch2)     //这里读取的顺序颠倒

        fmt.Println(<-ch1)

    }




func Right() {
    ch := make(chan string)
    go func(){
        // 在执行到这一步的时候main goroutine才会停止阻塞 子程序发送 主程序接收
        str := <- ch
        fmt.Println("receive data：" + str)
    }()
    ch <- "hello"
    fmt.Println("channel has send data")
}

/*
带缓冲的channel

带缓冲的channel的创建和不带缓冲的channel（也就是上面用的channel）的创建差不多，只是在make函数的第二个参数指定缓冲的大小。
// 创建一个容量为10的channel
ch := make(chan int, 10)

带缓冲的channel就像一个队列，遵从先进先从的原则，发送数据向队列尾部添加数据，从头部接受数据。
goroutine向channel发送数据的时候如果缓冲还没满，那么该goroutine就不会阻塞。
ch := make(chan int, 2)
// 前面两次发送数据不会阻塞，因为缓冲还没满
ch <- 1
ch <- 2
// goroutine会在这里阻塞
ch <- 3

反之如果接受该channel数据的时候，如果缓冲有数据，那么该goroutine就不会阻塞。

channel与goroutine之间的应用可以想象成某个工厂的流水线工作，流水线上面有打磨，上色两个步骤（两个goroutine），负责打磨的工人生产完成后会传给负责上色的工人，上色的生产依赖于打磨，两个步骤之间的可能存在存放槽（channel），如果存放槽存满了，打磨工人就不能继续向存放槽当中存放产品，直到上色工人拿走产品，反之上色工人如果把存放槽中的产品都上色完毕，那么他就只能等待新的产品投放到存放槽中。

备注

其实在实际应用中，带缓冲的channel用的并不多，继续拿刚才的流水线来做案例，如果打磨工人生产速度比上色工人工作速度要快，那么即便再多容量的channel，也会迟早被填满然后打磨工人会被阻塞，反之如果上色工人生产速度大于打磨工人速度，那么有缓冲的channel也是一直处于没有数据，上色工人很容易长时间处于阻塞的状态。

因此比较好的解决方法还是针对生产速度较慢的一方多加人手，也就是多开几个goroutine来进行处理，有缓冲的channel最好用处只是拿来防止goroutine的完成时间有一定的波动，需要把结果缓冲起来，以平衡整体channel通信。

*/

//单方向的channel
/*
使用channel来使不同的goroutine去进行通信，很多时候都和消费者生产者模式很相似，一个goroutine生产的结果都用channel传送给另一个goroutine，一个goroutine的执行依赖与另一个goroutine的结果。
 因此很多情况下，channel都是单方向的，在go里面可以把一个无方向的channel转换为只接受或者只发送的channel，但是却不能反过来把接受或发送的channel转换为无方向的channel，适当地把channel改成单方向，可以达到程序强约束的做法，类似于下面例子：
*/

func Single(){
    
    ch := make(chan string)
    go func(out chan<- string){
        out <- "hello"
    }(ch)

    go func(in <-chan string){
        fmt.Println(in)
    }(ch)

    time.Sleep(2 * time.Second)
}

/*
select多路复用

在一个goroutine里面，对channel的操作很可能导致我们当前的goroutine阻塞，而我们之后的操作都进行不了。而如果我们又需要在当前channel阻塞进行其他操作，如操作其他channel或直接跳过阻塞，可以通过select来达到多个channel（可同时接受和发送）复用。如下面我们的程序需要同时监听多个频道的信息：
broadcaster1 := make(chan string) // 频道1
broadcaster2 := make(chan string) // 频道2
select {
    case mess1 := <-broadcaster1:
        fmt.Println("来自频道1的消息：" + mess1)
    case mess2 := <-broadcaster2:
        fmt.Println("来自频道2的消息：" + mess2)
    default:
        fmt.Println("暂时没有任何频道的消息，请稍后再来~")
        time.Sleep(2 * time.Second)
}


select和switch语句有点相似，找到匹配的case执行对应的语句块，但是如果有两个或以上匹配的case语句，那么则会随机选择一个执行，如果都不匹配就会执行default语句块（如果含有default的部分的话）。
值得注意的是，select一般配合for循环来达到不断轮询管道的效果，可能很多小伙伴想着写个在某个case里用break来跳出for循环，这是不行的，因为break只会退出当前case，需要使用return来跳出函数或者弄个标志位标记退出

var flag = 0
for {
	if flag == 1 {break}
	select {
		case message := <- user.RecMess :
			event := gjson.Get(string(message), "event").String()
			if event == "login" {
				Login(message, user)
			}
			break
		case <- user.End :
			flag = 1
			break
	}
}
*/

// 关闭channnel
/*
channel可以接受和发送数据，也可以被关闭。
close(ch)


关闭channel后，所有向channel发送数据的操作都会引起panic，而被close之后的channel仍然可以接受之前已经发送成功的channel数据，如果数据全部接受完毕，那么再从channel里面接受数据只会接收到零值得数据。

channel的关闭可以用来操作其他goroutine退出，在运行机制方面，goroutine只有在自身所在函数运行完毕，或者主函数运行完毕才会打断，所以我们可以利用channel的关闭作为程序运行入口的一个标志位，如果channel关闭则停止运行。

无法直接让一个goroutine直接停止另一个goroutine，但可以使用通信的方法让一个goroutine停止另一个goroutine，如下例子就是程序一边运行，一边监听用户的输入，如果用户回车，则退出程序。

*/

func main0() {
    shutdown := make(chan struct{})
    var n sync.WaitGroup
    n.Add(1)
    go Running(shutdown, &n)
    n.Add(1)
    go ListenStop(shutdown, &n)
    n.Wait()
}

func Running(shutdown <-chan struct{}, n *sync.WaitGroup) {
    defer n.Done()
    for {
        select {
        case <-shutdown:
            // 一旦关闭channel，则可以接收到nil。 一旦nil值 就return跳出循环
            fmt.Println("shutdown goroutine")
            return
        default:
            fmt.Println("I am running")
            time.Sleep(1 * time.Second)
        }
    }
}

func ListenStop(shutdown chan<- struct{}, n *sync.WaitGroup) {
    defer n.Done()
    os.Stdin.Read(make([]byte, 1))
    // 如果用户输入了回车则退出关闭channel
    close(shutdown)
}

// 超时
func AfterTime() {
    ch := make(chan int)
    quit := make(chan bool)

	go func() {

		for {
			select {
				case num:=<-ch:
					fmt.Println("num = ", num)
				case <-time.After(time.Second * 3):
					fmt.Println("time out")
					quit <- true
					//return		// Goexit(）  等价
					goto ABC		// lable:标签名， 改名字任意
			}
		}
	ABC:		// 不能超出当前函数。
		fmt.Println("----break to lable----")
	}()

	for i:=0; i<2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit		// 阻塞
	fmt.Println("主 go 程 执行结束。")

}

//For-Range
/*
for-range语法可以用到通道上。循环会一直接收channel里面的数据，直到channel关闭。不同于array/slice/map上的for-range，channel的for-range只允许有一个变量。
for v = range aChannel {
	// use v
}

等价于
for {
	v, ok = <-aChannel
	if !ok {
		break
	}
	// use v
}
*/

// 使用多核

//Golang默认情况下都是使用一个cpu来执行goroutine的任务，所以在默认的情况下并不能执行并发任务。
//如果想使用多核并行的任务,可以通过runtime.GOMAXPROCS(）来设置CPU的个数的个数，当然这个数不能超过计算机拥有的CPU 。

func DoTask(wg *sync.WaitGroup) int {
    n := 2
    for i := 0; i < 20000; i++ {
        for j := 0; j < 100000; j++ {
            if n > 1000000 {
                n = n - 10000000
            } else {
                n++ 
            }   
        }   
    }   
    (*wg).Done()
    return n
}
func DoTasks(x int) {
    runtime.GOMAXPROCS(x)
    var wg sync.WaitGroup
    start := time.Now().UnixNano()
    for i := 0; i < 12; i++ {
        wg.Add(1)
        go DoTask(&wg)
    }   
    
    wg.Wait()
    fmt.Println("cpu", x, time.Now().UnixNano()-start, "ns")
}
func mainx() {
    for i := 1; i <= 8; i++ {
        DoTasks(i)
    }   
}    


// 利用通道来阻塞主程序
func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch: // 打印一个 fibonacci 数值
			fmt.Print(num, " ")
		case <-quit: // 接收到go程结束通知
			//return
			runtime.Goexit()
		}
	}
}
func main99() {
	ch := make(chan int)
	quit := make(chan bool)
	// 创建一个打印fibonacci数列的 子go程
	go fibonacci(ch, quit)
	x, y := 1, 1
	for i := 0; i < 50; i++ {
		ch <- x
		x, y = y, x+y
	}
	// 写完 20 个数，先阻塞主线程 等到子线程写入quit通道
	quit <- false
}

// runtime.Goexit()  // 将 go 程 结束
// runtime.Gosched()		// 主动让出 cpu 使用权



var wg sync.WaitGroup
var once sync.Once
 
func producer2(chanel1 chan int) {
    defer wg.Done()
    for i := 0; i < 100; i++ {
        chanel1 <- i
    }
    close(chanel1)
 
}
 
func consumer(ch1 chan int, ch2 chan int) {
    defer wg.Done()
    for v := range ch1 {
        ch2 <- v * v
    }
    //确保某个close操作被gorutines抢到后只被 close 1次
    once.Do(func() { close(ch2) })
 
}
 
func main67() {
    wg.Add(6)
    ch1 := make(chan int, 100)
    ch2 := make(chan int, 100)
    go producer2(ch1)
    go consumer(ch1, ch2)
    go consumer(ch1, ch2)
    go consumer(ch1, ch2)
    go consumer(ch1, ch2)
    go consumer(ch1, ch2)
    wg.Wait()
    for v := range ch2 {
        fmt.Println(v)
 
    }
 
}
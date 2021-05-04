package goroutine

import (
  "fmt"
  "time"
  "sync"
)

/*
并发安全 就是多个并发体在同一段时间内访问同一个共享数据，共享数据能被正确处理。

并发不安全的后果
例子
var wg sync.WaitGroup

//定义1个全局变量
var number int64

//对全局变量进行+1操作
func add1() {
    for i := 0; i < 5000; i++ {
        //1.从内存中找到number变量对应的值
        //2.进行+1操作
        //3.把结果赋值给number写到内存
        number++
    }
    wg.Done()
}
func main() {
    wg.Add(2)
    go add1()
    go add1()
    //fmt.Println(number)
    wg.Wait()
    fmt.Println(number) //每次执行结果都不一致
}

并发不安全最典型的案例就是卖票超售，设想有一家电影院，有两个售票窗口，售票员售票时候先看一下当前剩余票数是否大于0，如果大于0则售出票。
 用伪代码就是如下：
# 售票操作（一张票）
# 如果票数大于0
totalNum = getTotalNum()
if totalNum > 0
    # 则售出一张票
    totalNum = totalNum - 1
else
    failedToSold()

此时票数剩下一张票，两个售票窗口同时来了顾客，两个售票人都看了一下剩余票数还有一张，不约而同地收下顾客的钱，余票还剩一张，
但是却售出了两张票，就会出现致命的问题。

如何做到并发安全

目前最最主流的办法就是加锁就行操作，其实售票的整个操作同时间内只能一个人进行，在我看来归根到底加锁其实就是让查询和售票两个步骤原子化，只能一块执行，不能被其他程序中断，让这步操作变成串行化。

锁

锁的做法就是每次进入这段变量共享的程序片段，都要先获取一下锁，如果获取成功则可以继续执行，如果获取失败则阻塞，直到其他并发体把锁给释放，程序得到执行调度才可以执行下去。
 锁本质上就是让并发体创建一个程序临界区，临界区一次只能进去一个并发体，伪代码示意如下：
lock()
totalNum = getTotalNum()
if totalNum > 0
    # 则售出一张票
    totalNum = totalNum - 1
else
    failedToSold()
unlock()

读锁与写锁

读锁也叫共享锁，写锁也叫排它锁，锁的概念被发明了之后，人们就想着如果我很多个并发体大部分时间都是读，如果就把变量读取的时候也要建立临界区，那就有点太大题小做了。
于是人们发明了读锁，一个临界区如果加上了读锁，其他并发体执行到相同的临界区都可以加上读锁，执行下去，但不能加上写锁。这样就保证了可以多个并发体并发读取而又不会互相干扰。
队列

队列也是解决并发不安全的做法。多个并发体去获取队列里的元素，然后进行处理，这种做法和上锁其实大同小异，本质都是把并发的操作串行化，同一个数据同一个时刻只能交给一个并发体去处理。
 伪代码：
# 第一个获取到队列的元素就可以进行下去
isCanSold = canSoldList.pop()
totalNum = getTotalNum()
if totalNum > 0
    # 则售出一张票
    totalNum = totalNum - 1
else
    failedToSold()
*/



//定义1个全局变量
var x int64

//对全局变量进行+1操作
func add1(wg *sync.WaitGroup) {
    for i := 0; i < 5000; i++ {
        //1.从内存中找到number变量对应的值
        //2.进行+1操作
        //3.把结果赋值给number写到内存
        x++
    }
    wg.Done()
}


func main55() {
    var wg sync.WaitGroup
    wg.Add(2)
    //开启2个gorutines同时对x+1
    go add1(&wg)
    go add1(&wg)
    /*2个gorutines如果同1时刻都去获取公共变量x=50，
    然后在独自的栈中对x+1改变了x都=51
    就少+了1次，导致结果计算不准！
    */
    wg.Wait()
    fmt.Println(x) //每次执行结果都不一致
}

func add2(wg *sync.WaitGroup, lock *sync.Mutex) {
    defer wg.Done()
    for i := 0; i < 5000; i++ {
        lock.Lock()   //加锁
        defer lock.Unlock() //释放锁
        x++           //操作同1资源
        
    }
 
}

func main56() {
    var wg sync.WaitGroup
    var lock sync.Mutex
    wg.Add(2)
    //开启2个gorutines同时对x+1
    go add2(&wg,&lock)
    go add2(&wg,&lock)
    wg.Wait()
    fmt.Println(x) //每次执行结果都不一致
}


// 互斥锁
// 互斥锁需要确保的是某段时间内，不能有多个协程同时访问一段代码（临界区）。
type Bank struct {
 	sync.Mutex //继承锁
 	saving map[string]int // 每账户的存款金额
 }


func NewBank() *Bank {
 	b := &Bank{
 		saving: make(map[string]int),
 	}
 	return b
 }


 // Deposit 存款
func (b *Bank) Deposit(name string, amount int) {
 	b.Lock()
 	defer b.Unlock()


 	if _, ok := b.saving[name]; !ok {
 		b.saving[name] = 0
 	}
 	b.saving[name] += amount
 }


// Withdraw 取款，返回实际取到的金额
func (b *Bank) Withdraw(name string, amount int) int {
 	b.Lock()
 	defer b.Unlock()


 	if _, ok := b.saving[name]; !ok {
 		return 0
 	}
 	if b.saving[name] < amount {
 		amount = b.saving[name]
 	}
 	b.saving[name] -= amount


 	return amount
}


// Query 查询余额
func (b *Bank) Query(name string) int {
 	b.Lock()
 	defer b.Unlock()


 	if _, ok := b.saving[name]; !ok {
 		return 0
 	}


 	return b.saving[name]
}


func main1() {
 	b := NewBank()
 	go b.Deposit("xiaoming", 100)
 	go b.Withdraw("xiaoming", 20)
 	go b.Deposit("xiaogang", 2000)


 	time.Sleep(time.Second)
 	fmt.Printf("xiaoming has: %d\n", b.Query("xiaoming"))
 	fmt.Printf("xiaogang has: %d\n", b.Query("xiaogang"))
 }

// 读写锁
/*
读写锁要达到的效果是同一时间可以允许多个协程读数据，但只能有且只有1个协程写数据。

也就是说，读和写是互斥的，写和写也是互斥的，但读和读并不互斥。具体讲，当有至少1个协程读时，
如果需要进行写，就必须等待所有已经在读的协程结束读操作，写操作的协程才获得锁进行写数据。
当写数据的协程已经在进行时，有其他协程需要进行读或者写，就必须等待已经在写的协程结束写操作。

•Lock()和Unlock()是给写操作用的。
•RLock()和RUnlock()是给读操作用的。

上面的银行实现不合理：大家都是拿手机APP查余额，可以同时几个人一起查呀，这根本不影响，
银行的锁可以换成读写锁。存、取钱是写操作，查询金额是读操作，代码修改如下，其他不变：
*/

type Bank2 struct {
    sync.RWMutex
    saving map[string]int // 每账户的存款金额
}

func NewBank2() *Bank {
    b := &Bank{
        saving: make(map[string]int),
    }
    return b
}


// Deposit 存款
func (b *Bank2) Deposit(name string, amount int) {
    b.Lock()
    defer b.Unlock()


    if _, ok := b.saving[name]; !ok {
        b.saving[name] = 0
    }
    b.saving[name] += amount
}


// Withdraw 取款，返回实际取到的金额
func (b *Bank2) Withdraw(name string, amount int) int {
    b.Lock()
    defer b.Unlock()


    if _, ok := b.saving[name]; !ok {
        return 0
    }
    if b.saving[name] < amount {
        amount = b.saving[name]
    }
    b.saving[name] -= amount


    return amount
}


// Query 查询余额
func (b *Bank2) Query(name string) int {
    b.RLock()
    defer b.RUnlock()
    if _, ok := b.saving[name]; !ok {
        return 0
    }

    return b.saving[name]
}

func main2() {
    b := NewBank2()
    go b.Deposit("xiaoming", 100)
    go b.Withdraw("xiaoming", 20)
    go b.Deposit("xiaogang", 2000)

    time.Sleep(time.Second)
    print := func(name string) {
        fmt.Printf("%s has: %d\n", name, b.Query(name))
    }

    nameList := []string{"xiaoming", "xiaogang", "xiaohong", "xiaozhang"}
    for _, name := range nameList {
        go print(name)
    }

    time.Sleep(time.Second)
}


// 单次执行

//在程序执行前，通常需要做一些初始化操作，但触发初始化操作的地方是有多处的，但是这个初始化又只能执行1次，怎么办呢？
func main3() {
    var once sync.Once
    onceBody := func() {
        fmt.Println("Only once")
    }
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i := 0; i < 10; i++ {
        <-done
    }
}

/*
结果：
1➜  sync_pkg git:(master) ✗ go run once.go
2Only once  只执行了一次 调用了10次

*/

/*
sync.Map
var syncMap sync.Map
//新增
syncMap.Store(key, n)
//删除
syncMap.Delete(key)
//改
syncMap.LoadOrStore（key）
//遍历
syncMap.Range(walk)


golang中的map在并发情况下： 只读是线程安全的，但是写线程不安全，所以为了并发安全 & 高效，官方帮我们实现了另1个sync.map。

 
fatal error: concurrent map writes  //go内置的map只能支持20个并发写！
package main
 
import (
    "fmt"
    "strconv"
    "sync"
)
 
var m = make(map[string]int)
 
func get(key string) int {
    return m[key]
}
 
func set(key string, value int) {
    m[key] = value
}
 
func main() {
    wg := sync.WaitGroup{}
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func(n int) {
            key := strconv.Itoa(n)
            //设置1个值
            set(key, n)
            //获取1个值
            fmt.Printf("k=:%v,v:=%v\n", key, get(key))
            wg.Done()
        }(i)
    }
    wg.Wait()
}
　　

就支持20个并发也太少了！

Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。

同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

 

package main
 
import (
    "fmt"
    "strconv"
    "sync"
)
 
var syncMap sync.Map
var wg sync.WaitGroup
 
 
 
func walk(key, value interface{}) bool {
    fmt.Println("即将删除Key =", key, "Value =", value)
    syncMap.Delete(key)
    return true
}
 
func main() {
    for i := 0; i < 200; i++ {
        //开启20个协程去syncMap并发写操作,也是可以顺利写进去的的！
        key := strconv.Itoa(i)
        wg.Add(1)
        go func(n int) {
            //设置key
            syncMap.Store(key, n)
            //通过key获取value
            value, ok := syncMap.Load(key)
            if !ok {
                fmt.Println("没有该key", key)
            }
            fmt.Println(value)
            wg.Done()
        }(i)
 
    }
    //使用for 循环或者 for range 循环无法遍历所有syncMap只能使用syncMap.Range()
    //不幸运的Go没有提供sync.Map的Length的方法，需要自己实现！！
    syncMap.Range(walk)
    wg.Wait()
}



代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。

针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。

Go语言中原子操作由内置的标准库sync/atomic提供。



package main
 
import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)
 
type Counter interface {
    Inc()
    Load() int64
}
 
// 普通版
type CommonCounter struct {
    counter int64
}
 
func (c CommonCounter) Inc() {
    c.counter++
}
 
func (c CommonCounter) Load() int64 {
    return c.counter
}
 
// 互斥锁版
type MutexCounter struct {
    counter int64
    lock    sync.Mutex
}
 
func (m *MutexCounter) Inc() {
    m.lock.Lock()
    defer m.lock.Unlock()
    m.counter++
}
 
func (m *MutexCounter) Load() int64 {
    m.lock.Lock()
    defer m.lock.Unlock()
    return m.counter
}
 
// 原子操作版
type AtomicCounter struct {
    counter int64
}
 
func (a *AtomicCounter) Inc() {
    atomic.AddInt64(&a.counter, 1)
}
 
func (a *AtomicCounter) Load() int64 {
    return atomic.LoadInt64(&a.counter)
}
 
func test(c Counter) {
    var wg sync.WaitGroup
    start := time.Now()
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            c.Inc()
            wg.Done()
        }()
    }
    wg.Wait()
    end := time.Now()
    fmt.Println(c.Load(), end.Sub(start))
}
 
func main() {
    c1 := CommonCounter{} // 非并发安全
    test(c1)
    c2 := MutexCounter{} // 使用互斥锁实现并发安全
    test(&c2)
    c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
    test(&c3)
}
*/


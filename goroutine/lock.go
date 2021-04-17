package goroutine

import (
  "fmt"
  "time"
  "sync"
)

/*
并发安全 就是多个并发体在同一段时间内访问同一个共享数据，共享数据能被正确处理。

并发不安全的后果

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


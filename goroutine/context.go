package goroutine


import (
    "context"
    "fmt"
    "time"
)


// context
// context的本质是利用channel的close机制来触发的通知机制
// context是goruntine的上下文，包含goruntine的运行状态，环境，现场等信息。
// context在goruntine之间传递上下文信息，包括取消信号，超时时间，截止时间，传递参数等
// 不要讲context放在结构体中，直接将context类型放在函数的参数中的第一个位置，一般取名为ctx
// 不要向函数传递一个nil的context，可以使用TODO()
// 不要把本应该最为参数的值塞入context中，context存储的应该是一些公共的数据，如session，cookie等
// 同一个context可能被传到多个goruntine中，不要担心，context是并发安全的
/*
context源码
type Context interface {
    // 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
    Done() <-chan struct{}
    // 在 channel Done 关闭后，返回 context 取消原因
    Err() error
    // 返回 context 是否会被取消以及自动取消时间（即 deadline）
    Deadline() (deadline time.Time, ok bool)
    // 获取 key 对应的 value
    Value(key interface{}) interface{}
}
Done()返回一个只读channel，表示context被取消的信号，当这个channel被关闭时说明context被取消了

Err()返回一个错误，表示channel被关闭的原因，例如被取消还是已删除

Deadline() 返回context的截止时间

Value() 获取之前设置的key的value
*/

// func tree() {
	// 父context 可以 取消信号，超时时间，截止时间，传递参数 给 子context
// 	ctx1 := context.Background() // 根context
// 	ctx2, _ := context.WithCancel(ctx1)
// 	ctx3, _ := context.WithTimeout(ctx2, time.Second * 5)
// 	ctx5, _ := context.WithTimeout(ctx3, time.Second * 6)
// 	ctx6 := context.WithValue(ctx5, "userID", 12)
// 	fmt.Println(ctx6)
//   }

 
//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)
        select {
        case <-ctx.Done():
            fmt.Printf("done")
            return
        default:
            fmt.Printf("work")
        }
    }
}
 
func main_x0() {
	ctx, cancel := context.WithCancel(context.Background())
	// 5s后取消含ctx 的子协程
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go doStuff(ctx)
 
//10秒后取消doStuff
    time.Sleep(10 * time.Second)
    cancel() // 主进程传递取消信号
    fmt.Printf("down")
}


func doTimeOutStuff(ctx context.Context) {
    for {
        time.Sleep(1 * time.Second)
 
        if deadline, ok := ctx.Deadline(); ok { //设置了deadl
            fmt.Printf("deadline set")
            if time.Now().After(deadline) {
                fmt.Printf(ctx.Err().Error())
                return
            }
 
        }
 
        select {
        case <-ctx.Done():
            fmt.Printf("done")
            return
        default:
            fmt.Printf("work")
        }
    }
}
 
func main_x1() {
	// 5s超时
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    go doTimeOutStuff(ctx) 
    time.Sleep(10 * time.Second)
    cancel()
 
}


// 通过key value传递数据给含有ctx的子进程
func NewContext(ctx context.Context, userIP string) context.Context {
    return context.WithValue(ctx, "key", userIP)
}
 
// FromContext extracts the user IP address from ctx, if present.
func FromContext(ctx context.Context) (string, bool) {
    // ctx.Value returns nil if ctx has no value for the key;
    // the net.IP type assertion returns ok=false for nil.
    userIP, ok := ctx.Value("key").(string)
    return userIP, ok
}

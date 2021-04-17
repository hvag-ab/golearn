package goroutine

import (
	"fmt"
	"time"
	"sync"
)

// 每隔2s后=
func AfterTime2()  {
	// 打印当前时间
	fmt.Println("now:   ", time.Now())
	// 创建 Timer 定时器
	myTimer := time.NewTimer(time.Second * 2)
	for {
		time := <-myTimer.C		// 默认读——阻塞，定时时长2s到达后，系统写入当前时间，解除阻塞
		fmt.Println("定时后:", time)
	}
}

// 重置时间
func ResetTime()  {
	myTimer := time.NewTimer(time.Second * 5)

	go func() {
		<-myTimer.C
		fmt.Println("定时时间到！")
	}()
	//myTimer.Stop()
	myTimer.Reset(time.Second * 1)
	for {
		;
	}
}

// 2s后执行
func At()  {
	fmt.Println("now:   ", time.Now())

	time := <-time.After(time.Second * 2)

	fmt.Println("定时后:", time)
}


// 计时器
func Ti() {
    tick := time.Tick(1e8)
    boom := time.After(5e8)
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Println("    .")
            time.Sleep(5e7)
        }
    }
}

// 在11点35分打印"Golang"
func TimeEnd() {

	myTicker:=time.NewTicker(time.Second)		//设置时间周期
	for{
		nowTime:=<-myTicker.C		//当前时间
		if nowTime.Hour()==11 && nowTime.Minute()==35{
			fmt.Println("Golang")
			break
		}
	}
}

// 间隔执行
func Exec(){
	for range time.Tick(3 * time.Second) {
		fmt.Println("time")
	}
}


// 结束定时
func End(){

    ticker := time.NewTicker(5 * time.Second)
    quit := make(chan int)
    var wg  sync.WaitGroup
 
    wg.Add(1)
    go func(w *sync.WaitGroup) {
        defer wg.Done()
        fmt.Println("child goroutine bootstrap start")
        for {
            select {
                case <- ticker.C:
                    fmt.Println("ticker .")
                case <- quit:
                    fmt.Println("work well .")
                    ticker.Stop()
                    return
            }
        }
        // fmt.Println("child goroutine bootstrap end")
    }(&wg)
    time.Sleep(10 * time.Second)
    quit <- 1
	wg.Wait()
}
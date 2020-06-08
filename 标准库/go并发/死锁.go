package main

// func main() {
//     ch := make(chan int)
//     <- ch // 阻塞main goroutine, 通道被锁
// }

func main() {
    cha, chb := make(chan int), make(chan int)

    go func() {
        cha <- 1 // cha通道的数据没有被其他goroutine读取走，堵塞当前goroutine
        chb <- 0
    }()

    <- chb // chb 等待数据的写
}

// package main  加入缓冲可以解决死锁

// func main() {
//     cha, chb := make(chan int, 3), make(chan int)

//     go func() {
//         cha <- 1 // cha通道的数据没有被其他goroutine读取走，堵塞当前goroutine
//         chb <- 0
//     }()

//     <- chb // chb 等待数据的写
// }


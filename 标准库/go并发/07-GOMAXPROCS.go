package main

import (
	"fmt"
	"runtime"
)

func main() {
	_ = runtime.GOMAXPROCS(1) 	// 第一次 测试
	fmt.Printf("n = %d\n", 1)

	_ = runtime.GOMAXPROCS(2)         // 第二次 测试
	fmt.Printf("n = %d\n", 2)

	_ = runtime.GOMAXPROCS(8)         // 第二次 测试
	fmt.Printf("n = %d\n", 8)

	for {
		go fmt.Print(0)		// 子go程
		fmt.Print(1)			// 主go程
	}
}


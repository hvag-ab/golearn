/**
 *  datetime: 2017-05-23
 *  Go语言 - 递归
 *  Go中函数递归深度没有做限制
 */
package main

import (
	"fmt"
	"log"
)

func init(){//在main函数之前执行初始化函数  用来配置一些东西 比如日志
    log.SetPrefix("【Hvag】")
    log.SetFlags(log.LstdFlags | log.Lshortfile )
}


//1 * 2 * 3 * 4 * 5
func fact(n int) int {//阶乘
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fact(n-1)
}

func fib(n int) int{

	if n == 1 || n==2{
		return 1
	}else{
		return fib(n-1)+fib(n-2)
	}
}

func yh(n int) []int{
	// log.Println("start")
	if n==1{
		return []int{1}
	}
	if n==2{
		return []int{1,1}
	}else{
		pre := yh(n-1)
		length := len(pre)
		// log.Println(length)
		left := pre[:length-1]
		// log.Println(left)
		right := pre[1:]
		var after = []int{1}
		for i,_ := range left{
			// log.Println(i)
			after = append(after,left[i]+right[i])
		}
		after = append(after,1)
		return after
	}
}



func main() {

	// fmt.Println(fact(5))
	// fmt.Println(fib(35))

	for i:=1;i<100;i++{
		fmt.Println(yh(i))
	}
}

//$go run recursion.go

// 5
// 4
// 3
// 2
// 1
// 120

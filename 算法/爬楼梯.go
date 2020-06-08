package main

//爬楼梯 需要n楼才能到达楼顶，每次只能爬取1-2哥台阶 问有多少个方法

import "fmt"

type Stack struct {
	slc []int
}

func (s *Stack) Push(a int) {
	s.slc = append(s.slc, a)
}

func (s *Stack) Pop() int {
	a := s.slc[len(s.slc)-1]
	s.slc = s.slc[:len(s.slc)-1]
	return a
}
func (s *Stack) Len() int {
	return len(s.slc)
}

func ClimbStairs(n int) int {
	count := 0
	stack := &Stack{slc: []int{0}}   // 初始化
	for stack.Len() > 0 {
		current := stack.Pop()
		if current == n {
			count++
			continue
		}
		if current+1 <= n {
			stack.Push(current + 1)
		}
		if current+2 <= n {
			stack.Push(current + 2)
		}
	}
	return count
}

func main(){
	fmt.Println(ClimbStairs(10))
}

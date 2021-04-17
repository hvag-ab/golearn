package data_structure

import "fmt"
import "container/list"
import "container/heap"
import "container/ring"

func List(){
  l := list.New()
	l.PushBack(1) //尾插
	l.PushBack(2)
	print(l)

	fmt.Println("=========")

	l.PushFront(0) //头插
	print(l)

	fmt.Println("=========")

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			l.InsertAfter(1.1, e)
		}

		if e.Value == 2 {
			l.InsertBefore(1.2, e)
		}
	}

	print(l)

	fmt.Println("=========")

	fmt.Println(l.Front().Value) //返回链表的第一个元素
	fmt.Println("=========")

	fmt.Println(l.Back().Value) //返回链表的最后一个元素
	fmt.Println("=========")

	l.MoveToBack(l.Front())
	print(l)

	fmt.Println("=========")

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}


func Ring() {

	ring1 := ring.New(3)

	for i := 1; i <= 3; i++ {
		ring1.Value = i
		ring1 = ring1.Next()
	}

	ring2 := ring.New(3)

	for i := 4; i <= 6; i++ {
		ring2.Value = i
		ring2 = ring2.Next()
	}

	r := ring1.Link(ring2)

	fmt.Printf("ring length = %d\n", r.Len())

	r.Do(func(p interface{}) {
		fmt.Print(p.(int))
		fmt.Print(",")
	})

	fmt.Println()

	fmt.Printf("current ring is %v\n", r.Value)

	fmt.Printf("next ring is %v\n", r.Next().Value)

	fmt.Printf("prev ring is %v\n", r.Prev().Value)

	// ring 的遍历
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Print(p.Value.(int))
		fmt.Print(",")
	}

}

type Student struct {
  name  string
  score int
}

type StudentHeap []Student

func (h StudentHeap) Len() int { return len(h) }

func (h StudentHeap) Less(i, j int) bool {
  return h[i].score < h[j].score //最小堆
  //return stu[i].score > stu[j].score //最大堆
}

func (h StudentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StudentHeap) Push(x interface{}) {
  // Push and Pop use pointer receivers because they modify the slice's length,
  // not just its contents.
  *h = append(*h, x.(Student))
}

func (h *StudentHeap) Pop() interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0: n-1]
  return x
}

func Heap(){

      h := &StudentHeap{
        {name: "xiaoming", score: 82},
        {name: "xiaozhang", score: 88},
        {name: "laowang", score: 85}}

      // 初始化一个堆。一个堆在使用任何堆操作之前应先初始化。
      // Init函数对于堆的约束性是幂等的（多次执行无意义），并可能在任何时候堆的约束性被破坏时被调用。
      // 本函数复杂度为O(n)，其中n等于h.Len()。
      heap.Init(h)

      //向堆h中插入元素x，并保持堆的约束性。复杂度O(log(n))，其中n等于h.Len()。
      heap.Push(h, Student{name: "xiaoli", score: 66})

      for _, ele := range *h {
        fmt.Printf("student name %s,score %d\n", ele.name, ele.score)
      }

      for i, ele := range *h {
        if ele.name == "xiaozhang" {
          (*h)[i].score = 60

          // 在修改第i个元素后，调用本函数修复堆，比删除第i个元素后插入新元素更有效率。
          // 复杂度O(log(n))，其中n等于h.Len()。
          heap.Fix(h, i)
        }
      }

      fmt.Println("==========")

      for _, ele := range *h {
        fmt.Printf("student name %s,score %d\n", ele.name, ele.score)
      }

      fmt.Println("==========")

      for h.Len() > 0 {
        // 删除并返回堆h中的最小元素（取决于Less函数，最大堆或最小堆）（不影响堆de约束性）
        // 复杂度O(log(n))，其中n等于h.Len()。该函数等价于Remove(h, 0)
        item := heap.Pop(h).(Student)
        fmt.Printf("student name %s,score %d\n", item.name, item.score)
      }

    

}

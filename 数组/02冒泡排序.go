package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 5, 6, 3, 7, 10, 8, 2, 4}

	//冒泡排序

	for i := 0; i < 10-1; i++ {
		for j := 0; j < 10-1-i; j++ {
			if arr[j] > arr[j+1] {

				//数据交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//temp := arr[j]
				//arr[j] = arr[j+1]
				//arr[j+1] = temp
			}
		}
		fmt.Println(arr)
	}

	//fmt.Println(arr)
}

func Reversing(s []int64) []int64 {//切片反转
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
	  s[left], s[right] = s[right], s[left]
	}
	return s
  }

  //分块 主要用于当单个切片过大，需要分多次使用的时候，比如网络调用等。
func SliceChunk(s []int64, size int) [][]int64 {
	var ret [][]int64
	for size < len(s) {
		// s[:size:size] 表示 len 为 size，cap 也为 size，第二个冒号后的 size 表示 cap
		s, ret = s[size:], append(ret, s[:size:size])
	}
	ret = append(ret, s)
	return ret
}

func RemoveDuplicates(s []int64) []int64 {
	ret := s[:0]
	// 利用 struct{}{} 减少内存占用
	assist := map[int64]struct{}{}
	for _, v := range s {
	  if _, ok := assist[v]; !ok {
		assist[v] = struct{}{}
		ret = append(ret, v)
	  }
	}
	return ret
}

func FilterSlice(s []int64, filter func(x int64) bool) []int64 {
	// 返回的新切片
	// s[:0] 这种写法是创建了一个 len 为 0，cap 为 len(s) 即和原始切片最大容量一致的切片
	// 因为是过滤，所以新切片的元素总个数一定不大于比原始切片，这样做减少了切片扩容带来的影响
	// 同时，也有一个问题，因为 newS 和 s 共享底层数组，那么过滤后 s 也会被修改！
	newS := s[:0]
	// 遍历，对每个元素执行 filter，符合条件的加入新切片中
	for _, x := range s {
	  if !filter(x) {
		newS = append(newS, x)
	  }
	}
	return newS
}

func filter(x int) bool{
	if x==3{
		return true
	}
}
package main

import "fmt"

/*
把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：

传递数组的指针
使用数组的切片
*/

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // can also with dereferencing *a to get back to the array
		sum += v
	}
	return
}

// Output: The sum of the array is: 24.600000
//切片传递
// func sum(a []int) int {
//     s := 0
//     for i := 0; i < len(a); i++ {
//         s += a[i]
//     }
//     return s
// }

// func main {
//     var arr = [5]int{0, 1, 2, 3, 4}
//     sum(arr[:])
// }
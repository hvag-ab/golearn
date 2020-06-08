package main

import (
	"fmt"
	"strconv"
)
/*

*/

func main(){
	//str 转int
	i,_ := strconv.Atoi("3")
	// int - str
	i1 := strconv.Itoa(3)
	// parse 类函数
	i2,_ := strconv.ParseBool("true")
	i3,_ := strconv.ParseFloat("3.14234",64) // - float64 10代表十进制
	i4,_ := strconv.ParseInt("3.14234",10,64)
	i5,err := strconv.ParseUint("3.14234",10,64)
	fmt.Println(err)
	// Format 函数  就是转化为 str
	i6 := strconv.FormatInt(-42,16)
	// appengtp 

	//int - float
	i7 := float32(3)
	fmt.Println(i,i1,i2,i3,i4,i6,i5,i7)

}
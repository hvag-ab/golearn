package main



//defer 在函数返回值后在调用 调用完后 在return 所以 先返回x=10 在调用defer 最后在执行returnp0
//defer 用在关闭资源 比如close 文件 数据库 
func main0302() int {

	x := 10
	defer func(x int) {
		x++
	}(2)
	return x
}

func main03033() (y int) {

	x := 10
	defer func() {
		x++
	}()
	return x
}


func main(){
	println(main0302())
	println(main03033())
}

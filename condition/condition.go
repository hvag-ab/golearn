package condition

import "fmt"
import "time"

func If(){
	v1 := 0
	if v1 > 1 {
		fmt.Println("v1 > 1")
	} else if v1 == 0 {
		fmt.Println("v1 == 1")
	} else {//else不是换行 只能在前一}后面接着写
		fmt.Println("v1 < 1")
	}

	//支持初始化语句

	if v2 := 0; v2 < 10 {
		fmt.Println("v2 < 10")
	}

	// go 不支持三元操作符
}

func For(){
	// 最常见的for循环结构  - 切记这里的i的作用域只是在for循环内部
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	n :=5
	for ;n>0;n /=2{//多条件 注意分号 条件写到for表达式上 第一个分号前面省略了n的初始值
		fmt.Println("hvag",n)
	}

	//for 模拟while （golang中无while循环结构）
	s1 := 0
	for s1 < 10 {
		s1++
	}
	fmt.Println("s1 = ", s1)

	//死循环
	for { //for true
		i := 0
		i++
		fmt.Println(i)
	}

}

func Switch(){
	var m int

	fmt.Scan(&m)

	//switch 变量  case 值1:代码1 case 值2: 代码2 default:代码3
	switch m {
	case 1:
		fmt.Println("一月份")
	case 2:
		fmt.Println("二月份")
	default:
		fmt.Println("未知月份")
	}

	a := 10
	b := 20

	//if 优点可以判断区间 可以进行嵌套使用  缺点 执行效率比较低
	//if a > b {
	//	fmt.Println("a")
	//} else {
	//	fmt.Println("b")
	//}

	//switch 语句 执行效率高  缺点 不能判断复杂区间 不能嵌套使用
	switch a > b {

	case true:
		fmt.Println(a)
	case false:
		fmt.Println(b)

	}

	//case 分支  支持多个值

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// case 分支 可以是表达式。可代替if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's Before noon")
	default:
		fmt.Println("It's After noon")
	}


}

func Break(){
	sum:=0
	i:=0
	for  {
		if i %2 == 0{
			continue
		}else{
			sum+=i
		}
		if i==100{
			break
		} 
		i++
	}


	/*
		goto语句可以无条件地转移到过程中指定的行。
	通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
	在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱
	goto对应(标签)既可以定义在for循环前面,也可以定义在for循环后面，当跳转到标签地方时，继续执行标签下面的代码。
		*/
	//  放在for前面，此例会一直循环下去
	//    Loop://标识 变量可以自己定义
	//    fmt.Println("test")
	for a:=1;a<5;a++{
		fmt.Println(a)
		if a>3{
			goto Loop //goto表示跳转到哪一行运行 使用标识
		}
	}

	Loop://放在for后边 可前可后
	fmt.Println("test")


}
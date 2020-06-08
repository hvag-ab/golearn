package main

import "fmt"
import "os"

func main0501() {
	/*
	Print:   输出到控制台(不接受任何格式化，它等价于对每一个操作数都应用 %v)
    Println: 输出到控制台并换行,可以打印出字符串，和变量 
     Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出整形变量和整形）
Sprintf：格式化并返回一个字符串而不带任何输出。
         s := fmt.Sprintf("a %s", "string") fmt.Printf(s)
Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。
         fmt.Fprintf(os.Stderr, “an %s\n”, “error”)
	*/

	a := 10
	fmt.Println(a)//right
	fmt.Println("abc")//right
	fmt.Printf("%d",a)//right
	// fmt.Printf(a)//error
	fmt.Print("gggg")
	fmt.Print(a)//输出error
}
func main0502(){
	a:=10
	//\n 是一个转义字符 表示换行
	fmt.Printf("%T\n",a)
	//%d是一个占位符 表示输出一个10进制整型数据
	fmt.Printf("%d\n",a)
	b:=3.595
	//%f是一个占位符 表示输出一个浮点型数据 默认小数位数保留6位
	fmt.Printf("%f\n",b)
	//%.2f表示数据保留两位小数  会对第三位小数四舍五入
	fmt.Printf("%.2f\n",b)

	c:="hello world"
	fmt.Printf("%T\n",c)//string  字符串类型
	//%s是一个占位符表示输出一个字符串类型数据
	fmt.Printf("%s\n",c)

	d:='a'
	fmt.Printf("%T\n",d)//byte 字符类型  int32  ASCII 对应整数的值
	//%c是一个占位符表示输出一个字符类型数据
	fmt.Printf("%c\n",d)

	e:=true//false true
	fmt.Printf("%T\n",e)//bool 布尔类型 用作于条件判断
	//%t是一个占位符表示输出一个bool类型数据
	fmt.Printf("%t\n",e)
}

func main0601() {
	var score int
	//通过键盘为score赋值  & 取地址运算符
	//fmt.Scan(&score)
	fmt.Scanf("%d", &score)
	fmt.Println(score)
}
func mainxx() {
	//通过键盘输入三门成绩  计算总成绩和平均成绩
	var c, m, e int

	//通过键盘为三门成绩赋值  数据间隔空格 或者回车
	//fmt.Scan(&c, &m, &e)
	fmt.Scanf("%d%d%d", &c, &m, &e)
	//计算总成绩
	sum := c + m + e

	fmt.Println("总成绩：", sum)
	fmt.Println("平均成绩： ", sum/3) //整型数据相除 得到的结果是整型数据



}

type point struct {
    x, y int
}


func main(){

    //Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
    p := point{1, 2}
    fmt.Printf("%v\n", p) // {1 2}
    //如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
    fmt.Printf("%+v\n", p) // {x:1 y:2}
    //%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
    fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
    //需要打印值的类型，使用 %T。
    fmt.Printf("%T\n", p) // main.point
    //格式化布尔值是简单的。
    fmt.Printf("%t\n", true)
    //格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
    fmt.Printf("%d\n", 123)
    //这个输出二进制表示形式。
    fmt.Printf("%b\n", 14)
    //这个输出给定整数的对应字符。
    fmt.Printf("%c\n", 33)
    //%x 提供十六进制编码。
    fmt.Printf("%x\n", 456)
    //对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
    fmt.Printf("%f\n", 78.9)
    //%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
    //使用 %s 进行基本的字符串输出。
    fmt.Printf("%s\n", "\"string\"")
    //像 Go 源代码中那样带有双引号的输出，使用 %q。
    fmt.Printf("%q\n", "\"string\"")
    //和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
    fmt.Printf("%x\n", "hex this")
    //要输出一个指针的值，使用 %p。
    fmt.Printf("%p\n", &p)
    //当输出数字的时候，你将经常想要控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    //你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    //要最对齐，使用 - 标志。
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
    //你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    //要左对齐，和数字一样，使用 - 标志。
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
    //到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。Sprintf 则格式化并返回一个字符串而不带任何输出。
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
    //你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	
	main0501()
}
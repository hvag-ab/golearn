package basic

import (
	."fmt"
	"math"
	"reflect"
	"strings"
	"unsafe"
)

// 字符串类型 string 
func Str(){
	var str string = "hvag"

	//strings.Contains 可以用于模糊查找 相当于python的in
	str1 := strings.Contains(str,"hv")

	str2 := []string{"a", "b", "c", "d"}
	//字符串拼接  可以将一个字符串切片组合成一个字符串，可以使用连接符
	s := strings.Join(str2, "-")

	// 查找一个字符串在另一个字符串中第一次出现的位置 返回值为整型 数据下标 如果没找到返回值为-1
	i := strings.Index(str, "h")

	// 字符串切片
	substr := str[:3]

	////将一个字符串中的字符串用另外一个字符串替换  替换n次  -1表示全部替换
	str3:=strings.Replace(str,"h","m",-1)
	// split
	str4:=strings.Split(str,"h")

	//去掉头尾指定字符串
	str5:=strings.Trim(str,"g")

	//去掉字符串中的空格 并转成切片类型
	str6:=strings.Fields(str)

	//字符串原型格式
	var strFormat = `
		hello
			world
	`
	Println(strFormat)

	Println(str,str1,str2,str3,str4,str5,str6,s,i,substr)

}

// byte 类型 
func Byte(){
		/**
	 *  byte
	 *  宽度：1字节
	 *  默认值：0
	 *  字节类型，uint8别名，可以看作为由8位二进制标示的无符号整数类型
	 *  范围：0 ~ 255
	 *  注意用单引号  字符串用双引号
	 */
	 var firstByte byte
	 firstByte = 'a'                               // firstByte > 255 or firstByte < 0  is error
	 Println(firstByte, reflect.TypeOf(firstByte)) //0, uint8
}

// bool 
func Bool() {
	/*
	 *  布尔类型 false true
	 *	宽度：1字节
	 *  默认值：false
	 */
	 var isActive bool
	 Println(isActive, reflect.TypeOf(isActive))
 
}

// int int64 int32 uint uint32 uint64 rune(int32别名, 专注于存储unicode编码的单个字符)
func Integer(){
	var firstRune rune
	firstRune = 'a'
	firstRune = 29999
	Println(firstRune, unsafe.Sizeof(firstRune)) //占用4个字节

	var (
		a int = 32
		// b int32
		// c int64
		// d uint 
		// e uint32
		// f uint64
	)
	Println(reflect.TypeOf(a), unsafe.Sizeof(a)) // int 8
}

// float32 float64
func Float(){
	/**
	 *  float32, float64
	 *  浮点型
	 *  默认值： 0
	 *  字节：4字 / 8字节
	 */

	 var f1 float32
	 Println(f1)
	 var f2 float64
	 Println(f2)

	//可以通过标准库math，查看各个数字类型的取值范围

	// math.MaxFloat32
	// math.MaxFloat64
	// math.MaxInt16
	// math.MaxInt32
	// math.MaxInt8
	// math.MaxUint16
	// math.MaxUint32
	// math.MaxUint64
	// math.MaxUint8
}

// 数据类型可用标识
func Num(){
	// 数值类型可用标识：8进制，16进制，科学计数法
	a, b, c, d := 071, 0x1F, 1e9, math.MaxUint16
	_, _, _, _ = a, b, c, d
}

// complex64, complex 128
func Complex(){
	/**
	 *   complex64, complex 128
	 *   复数类型
	 *   由float32 / float64 类型的实部和虚部联合表示
	 *   默认值：0+0i
	 */

	 var c1 complex64
	 c1 = 1 + 1i
	 Println(c1)
 
	 var c2 complex128
	 c2 = 1.0 + 3.8i
	 Println(c2)
}

// 运算符
func Opt(){
	a := 10
	b := 20

	// 赋值运算符 = += -= *= /= %=
	a *= b
	a += b
	a -= b
	a /= b//a=a/b
	a %= b //a=a%b

	// 逻辑运算符 && 逻辑and || 逻辑or  ！逻辑非
	d := a > b && a < a+2
	Println(d)
	// 比较运算符  > >= <= < ==  != 返回bool
	a1 := a>b
	b1 := a<=b
	c1 := a==b
	Println(a1,b1,c1)
	// 算数运算符 + - * /  ++ -- %
	c := a/b
	Println(a,b,c)
	// 整数相除得到的结果也是整数  0不能作为除数
	//自增和自减不能出现在表达式中  引起程序的二义性   //a=a++ - a-- * a++//err
	// a++
	// a--
}

// 常量
func Const(){
	/**
	*  datetime: 2017-05-21
	*  Go语言-常量
	*  Go语言的常量值只能包含：字符串值、布尔值、数值类型(int,float,complex等)
	*  常量的值还可以使用 len, cap, unsafe.Sizeof 等编译器内可确定结果的函数返回值
	*  常量命名规范：官方推荐使用驼峰方式命名(以大写开头的常量为导出常量，小写开头的为包常量)
	*  当使用常量组的时候，如果常量不提供类型和初始化值，那么该常量的值和类型视作和上一个常量相同
	*/
	// 声明ip常量，并且常量类型通过值推到.
	const ip string = "192.168.1.0"
	// 声明hvag常量
	const hvag = "hvag"
	const d = 1.3
	
	//声明一组常量
	const (
		c1, c2 = 10, 20
		c4     = false
	)
	
	//声明组常量
	const (
		c5 = 100
		c6 //c6的值为100。(如果常量不提供类型和初始化值，那么该常量的值和类型视作和上一个常量相同)
		c7 = 102
	)
	const (
		c8  = "c8 constant"
		c9  = len(c8)
		c10 = unsafe.Sizeof(c8)
		c11 = cap([3]int{1, 2, 3})
	)
	// 星期
	const (
		//星期天 初始化为0，后面为以后每一行为自增量
		Sunday = iota
		//星期一
		Monday
		//星期二
		Tuesday
		//星期三
		Wednesday
		//星期四
		Thurday
		//星期五
		Friday
		//星期六
		Saturday
	)
 
}

// print 格式化  输出
func Pprint(){
	/*
	Print:   输出到控制台(不接受任何格式化，它等价于对每一个操作数都应用 %v)
    Println: 输出到控制台并换行,可以打印出字符串，和变量 
    Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出整形变量和整形）
	Sprintf：格式化并返回一个字符串而不带任何输出。
         s := fmt.Sprintf("a %s", "string") fmt.Printf(s)
	Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout。
		 fmt.Fprintf(os.Stderr, “an %s\n”, “error”)
	//通过键盘为score赋值  & 取地址运算符
	//fmt.Scan(&score)	 
	fmt.Scanf("%d%d%d", &c, &m, &e)
	//计算总成绩
	
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
	   */

}

// 指针 
func Ptr(){
	var a int = 10

	Printf("%p\n", &a)

	//将a的地址赋值给一个指针变量
	//指针* 表示一级指针
	var p *int = &a //定义一个p变量 指向a int变量的内存地址
	Printf("%p\n", p)
	Printf("%p\n",&p)
	Printf("Here is the string *p: %d\n", *p) // prints string *p就取 a的值

	//通过指针间接修改变量的值
	*p = 100

	var p1 *int //空指针
	//fmt.Printf("%p\n",p)
	Println(p1)
	//指针变量指向了内存地址编号为0的空间  0-255为系统占用不允许读写操作
	// *p1=100//err //写入操作 必须要指向一个内存地址 否则报错
	// Println(*p1) //err 读操作

	//野指针   指针变量指向内存中一个未知的空间
	//从操作野指针对应的内存空间会报错
	//var p *int =*int(0xff00)
	// var a2 int = 10
	// var p2 *int = &a2
	//野指针
	//p2 = 100

	//new  创建一块新的空间 将空间的地址赋值给一个指针变量
	var p4 *int
	p4 = new(int) //new 新建内存默认存储的值和（）内的类型是一致的
	*p4 = 100
	Println(*p4)

	//多级指针
	a5 := 10 //int
	p5 := &a5 //*int

	//pp := &p5//**int
	//定义二级指针存储一级指针的地址
	var pp **int = &p5


	var p6 ***int = &pp

	**pp = 123
	//二级指针的值
	//*p3
	//一级指针的值
	//**p3
	//变量的值
	***p6=123

	Printf("%T\n", a5)
	Printf("%T\n", p5)
	Printf("%T\n", pp)

}

// math
func Math() {
    Println(math.Abs(float64(-3.2)))         //取到绝对值
    Println(math.Ceil(3.8))             //向上取整
    Println(math.Floor(3.6))             //向下取整
    Println(math.Mod(11,3))         //取余数 11%3 效果一样
    Println(math.Modf(3.22))             //取整数跟小数
    Println(math.Pow(3,2))             //X 的 Y次方  乘方
    Println(math.Pow10(3))             //10的N次方 乘方
    Println(math.Sqrt(9))             //开平方  3
    Println(math.Cbrt(8))             //开立方  2
    Println(math.Pi)                     //π
	Println(math.E)                     //e
    Println(math.Round(4.2))          //四舍五入
    Println(math.Max(-1.3, 0))     //0   返回x和y中最大值
    Println(math.Min(-1.3, 0))    //-1.3  返回x和y中最小值
	Println(math.Log(3))

}
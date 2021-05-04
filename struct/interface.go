package struct_

/**
 *  datetime: 2017-05-24
 *  Go语言-interface(接口类型)
 *  接口类型用于定义一组方法行为, 并且方法只有方法声明，没有方法体。
 *  一个接口类型声明，可以嵌套其他接口类型。
 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
 接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。

通过如下格式定义接口：

type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
上面的程序定义了一个结构体 Square 和一个接口 Shaper，接口有一个方法 Area()。

在 main() 方法中创建了一个 Square 的实例。在主程序外边定义了一个接收者类型是 Square 方法的 Area()，用来计算正方形的面积：结构体 Square 实现了接口 Shaper 。

所以可以将一个 Square 类型的变量赋值给一个接口类型的变量：areaIntf = sq1 。

现在接口变量包含一个指向 Square 变量的引用，通过它可以调用 Square 上的方法 Area()。当然也可以直接在Square 的实例上调用此方法，但是在接口实例上调用此方法更令人兴奋，它使此方法更具有一般性。接口变量里包含了接收者实例的值和指向对应方法表的指针。

这是 多态 的 Go 版本，多态是面向对象编程中一个广为人知的概念：根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为
如果 Square 没有实现 Area() 方法，编译器将会给出清晰的错误信息：

cannot use sq1 (type *Square) as type Shaper in assignment:
*Square does not implement Shaper (missing Area method)
*/

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func Interface() {
	sq1 := new(Square)
	sq1.side = 5

	// var areaIntf Shaper
	// areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	var areaIntf Shaper = sq1 //接口绑定哪一个结构体 便于判断接口属于哪一个类型
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	//第二种接口绑定实例
	var i Shaper
	var sq Square = Square{side: 12} //实例化
	i = &sq                          //结构i绑定在实例地址上
	fmt.Printf("The square has area: %f\n", i.Area())
}

//嵌套接口
type ReadWrite interface {
	Read() bool
	Write() bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

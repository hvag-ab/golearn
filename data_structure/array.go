package data_structure

import "fmt"

/*
在Go语言中数组是一个值类型（value type）。是真真实实的数组，而不是一个指向数组内存起始位置的指针，
也不能和同类型的指针进行转化，这一点严重不同于C语言。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。
如果将数组作为函数的参数类型，则在函数调用时该参数将发生数据复制。因此，在函数体中无法修改传入的数组的内容，
因为函数内操作的只是所传入数组的一个副本。
定义方式如下

var arr [n]type
    n表示数组长度，type表示数组存储类型。

   在Go语言中，数组长度在定义后就不可更改，在声明时长度可以为一个常量或者一个常量表达式（常量表达式是指在编译期即可计算结果的表达式）。数组的长度是该数组类型的一个内置常量，可以用Go语言的内置函数len()来获取。

arrLength := len(arr)

*/

func Array(){
    var arr1 [5]int
    arr2 := [5]int{1, 2, 3, 4, 5}   //指定长度为5，并赋5个初始值
    arr3 := [5]int{1, 2, 3}         //指定长度为5，对前3个元素进行赋值，其他元素为零值
    arr4 := [5]int{4: 1}            //指定长度为5，对第5个元素赋值
    arr5 := [...]int{1, 2, 3, 4, 5} //不指定长度，对数组赋以5个值
    arr6 := [...]int{8: 1}          //不指定长度，对第9个元素（下标为8）赋值1
    fmt.Println(len(arr3))
    fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
    // [0 0 0 0 0] [1 2 3 4 5] [1 2 3 0 0] [0 0 0 0 1] [1 2 3 4 5] [0 0 0 0 0 0 0 0 1]

    // for访问数组
    for i := 0; i < len(arr3); i++{
        fmt.Printf("arr[%d]=%d\n", i, arr3[i])
    }

    // range 访问数组
    for index, value := range(arr3) {
        fmt.Printf("arr[%d]=%d\n", index, value)
    }

    // 数组值传递 函数是值传递 所以相当于是深拷贝 修改不影响
    // 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
    // 传递数组的指针
    // 使用数组的切片

    // func mod(arr [5]int) {
    //     arr[0] = 10
    //     fmt.Println("In modify(), arr values:", arr)
    // }
    // arr := [5]int{1, 2, 3, 4, 5}
    // mod(arr)
    // fmt.Println("In main(), arr values:", arr)
    // In modify(), arr values: [10 2 3 4 5]
    //In main(), arr values: [1 2 3 4 5]

    	//二维数组
	//var 数组名 [行][列]数据类型
	//数组元素个数=行*列

	//var arr [3][4]int = [3][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}}
	//数组如果初始化部分值  未初始化值为0
	//var arr [3][4]int = [3][4]int{{1, 2, 3, 4}, {3}}
	//数组可以找到具体下标可以赋值
    //var arr[3][4]int=[3][4]int{1:{1,2,3,4},2:{3,4,5,6},0:{'a','b','c','d'}}
    
    // 数组指针
    var arr10 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//定义指针指向数组
	//p := &arr
	//定义指针是要和数组元素个数数据类型保持一致
	var p *[10]int = &arr10
	//数组首地址
	//fmt.Printf("%p\n",p)
	//fmt.Printf("%p\n",&arr)
	//fmt.Printf("%p\n",&arr[0])

	//fmt.Printf("%T\n", p)
	//通过指针间接修改数组元素的值
	(*p)[0]=123
	(*p)[1]=222
	// 数组指针加下标就可以直接修改数组元素的值
	p[0]=123
	p[1]=222
	//
	fmt.Println(arr10)

	//让数组指针和数组建立关系
	// for i := 0; i < len(p); i++ {
	// 	fmt.Println(p[i])
	// }

	for i, v := range p {
		fmt.Println(i, v)
    }
    
    // 指针数组
    a := 10
	b := 20
	c := 30
    var arr11 [3]*int
	arr11[0] = &a
	arr11[2] = &b
	arr11[1] = &c

    fmt.Println(arr11)
    
    //通过指针数组间接修改变量的值
	*arr11[1]=222
    *arr11[0]=123
    
    //遍历指针数组
	for i:=0;i<len(arr11);i++{
		fmt.Println(*arr11[i])
    }
    
    // p = new([10]int)//使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针

	// //当new([10]int)时会在内存中创建一个连续的整型空间10个 默认值为0  数组指针变量p指向该内存
	// for i := 0; i < len(p); i++ {
	// 	p[i] = i
	// }
	// //fmt.Println(*p)

	// for i, v := range p {
	// 	fmt.Println(i, v)
	// }


}

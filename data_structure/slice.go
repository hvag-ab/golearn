package data_structure

/*
slice 数据结构
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}


golang 源码
len <= cap
切片的零值是nil, 且长度(len)和容量(cap) 都是0
•基于数组或者slice生成一个slice的时候,新的slice和原来数组/slice 的底层数组是同一个
•基于数组或者slice生成slice的cap=原来对应的数组长度-现在slice的第一个元素对应的索引
•slice 作为参数传递的是 副本,但是对应的数组的指针不变
•扩容规则:
在一般的情况下，你可以简单地认为新切片的容量（以下简称新容量）将会是原切片容量（以下
 简称原容量）的 2 倍。
 但是，当原切片的长度（以下简称原长度）大于或等于1024时，Go 语言将会以原容量的1.25
倍作为新容量的基准（以下新容量基准）

slice 数组指针什么情况下会发生变化?

确切地说，一个切片的底层数组永远不会被替换。为什么？虽然在扩容的时候 Go 语言一定会生
 成新的底层数组，但是它也同时生成了新的切片。它是把新的切片作为了新底层数组的窗口，而
 没有对原切片及其底层数组做任何改动。
 实际测试:
扩容的时候slice的数组发生了变化,但是slice并没有发送变化.

//   变量名            变量类型
var variable_name = []var_type
*/
import "fmt"

func SliceBase(){
//   var sli []int // 如果切片只声明而没有初始化，那么这个切片的默认值为nil,长度为 0
    //通过new创建一个切片
	// p = new([]int)

	// fmt.Printf("%p\n", p)
	// *p = append(*p, 1, 2, 3, 4, 5)
	// *p = append(*p, 1, 2, 3)
	// (*p)[1] = 222

//   // 创建了一个类型为[]int,长度为10，容量为20的切片，如果不指定切片的容量，例如slice := make([]int, 10),那么该切片的容量和长度相等。
//   slice1 := make([]int, 10, 20)

  // 还有一种创建切片的方式，是使用字面量，就是指定初始化的值。
  slice2 := []int{1,2,3,4,5}

  // 这是指定了第5个元素为1，其他元素都是默认值0
//   slice3 := []int{4:1}

  // 数组赋值
//   the_array := [5]int {1, 2, 3, 4, 5}
//   the_slice := the_array[2:5]


  // 切片索引 类似python索引 注意也是从0开始
//   slice21 := slice2[:]
//   slice22 := slice2[0:]
//   slice23 := slice2[:5]

  // 指定修改某个值 赋值
  newSlice := slice2[1:3]
  newSlice[0] = 10
  // newSlice and slice2 同一个内存地址 所以当修改的时候，底层数组的值就会被改变，所以原切片的值也改变了
  fmt.Printf("%p\n", &slice2[1])
  fmt.Printf("%p\n", &newSlice[0])

  /*
  对于底层数组容量是k的切片slice[i:j]来说
  长度:j-i
  容量:k-i
  比如我们上面的例子slice[1:3],长度就是3-1=2，容量是5-1=4。不过代码中我们计算的时候不用这么麻烦，因为Go语言为我们提供了内置的len和cap函数来计算切片的长度和容量。
  */

  slice4 := []int{1, 2, 3, 4, 5}
  newSlice2 := slice4[1:3]
  fmt.Printf("newSlice2长度:%d,容量:%d",len(newSlice2),cap(newSlice2))

  // 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!
  arr := [...]int{1, 2, 3, 4 ,5}
  slice5 := []int{1, 2, 3, 4, 5}
  fmt.Printf("arr的首地址为: %p\n", &arr)
  fmt.Printf("slice的首地址为: %p\n", slice5)

  // 切片算是一个动态数组，所以它可以按需增长，我们使用内置append函数即可
  // append()，添加元素时，如果超越了cap，系统会自动分配一个更大cap的底层数组，初始的数组会被gc
  slice6 := []int{1, 2, 3, 4, 5}
  newSlice3 := slice6[1:3]

  newSlice3=append(newSlice3,10)
  newSlice3=append(newSlice3,10,20,30)

  // 迭代切片
   slice7 := []int{1, 2, 3, 4, 5}
   for i := 0; i < len(slice7); i++ {
       fmt.Printf("值:%d\n", slice7[i])
   }

   //在函数间传递切片 函数间的传递都是值传递 slice是指针 相当于传递的是内存地址 修改后内存地址不变 浅拷贝
//    func modify(slice []int) {
//     fmt.Printf("%p\n", slice)
//     slice[1] = 10
//     }

//     slice9 := []int{1, 2, 3, 4, 5}
//     fmt.Printf("%p\n", slice9)
//     modify(slice9)
//     fmt.Println(slice9)
    /*
    0xc0000180c0
    0xc0000180c0
    所以这两个切片指向同一个内存地址。因此我们修改一个索引的值后，发现原切片的值也被修改了，说明它们共用一个底层数组。
    */

    // copy
    //1.不同类型的切片无法复制
    //2.如果s1的长度大于s2的长度，将s2中对应位置上的值替换s1中对应位置的值
    //3.如果s1的长度小于s2的长度，多余的将不做替换
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    s3 := []int{6, 7, 8, 9}
    copy(s1, s2)
    fmt.Println(s1) //[4 5 3]
    copy(s2, s3)
    fmt.Println(s2) //[6 7]

    // 切片指针
    slice11 := []int{1, 2, 3, 4, 5}

	//指向切片的指针
	var p11 *[]int = &slice11
	//切片不允许直接通过指针来操作元素 err
	//	p11[0]=123
	//	p11[2]=222

	(*p11)[0] = 123
	(*p11)[2] = 222
	fmt.Println(slice11)


}

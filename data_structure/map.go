package data_structure

import "fmt"

/*
map是一堆未排序键值对的集合，map的key必须为可比较类型,比如 == 或 !=，map查找比线性查找快,但慢于索引查找(数组，切片)

定义一个map

格式 var name map[keytype]valuetype

其中的keytype为map中键的类型,valuetype为map中值的类型,map还可以使用 := 或者 make()创建,如:

*/
func Map(){
    // var ai map[int]int
    // var b = map[string]string{}
    // c := map[string]interface{}{}
    // d := make(map[string]int)

    // 初始化
    var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
    // 访问元素
    fmt.Println(a["name"], a["age"])
    // 在golang中 如果访问一个未定义的key 将返回这个map中value的默认值,来验证一下
    // fmt.Println(a["parent"]) //不存在 返回""

    // 格式 value,bool := map[key] 如下列程序
    if value, ok := a["parent"]; ok == true {
        fmt.Println(value)
    } else {
        fmt.Println("key not in map")
    }

    // update
    a["name"] = "李四"
    fmt.Println(a)

    // add 新增元素 map[key] = value 的方式增加元素
    a["hvag"] = "hvag"

    // 删除元素 格式 delete(map,key)
    delete(a, "hvag")
    fmt.Println(a)


   // map是引用类型
    var aa = map[int]int{1: 1, 2: 2}
    bi := aa
    aa[1] = 123456
    fmt.Println(aa, bi)
    // 以上程序会输出 map[1:123456 2:2] map[1:123456 2:2]

    // 函数  map传递的是内存地址值
    // func test(a map[int]int) {
    // a[1] = 111
    // }

    // var ac = map[int]int{1: 1, 2: 2}
    // fmt.Println(ac)
    // test(a)
    // fmt.Println(ac)
    // 以上程序会输出 map[2:2 1:111] map[1:111 2:2]

    // 遍历
    for key, value := range a {
        fmt.Println(key, value)
   }








}

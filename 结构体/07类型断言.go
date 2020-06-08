package main

import (
    "fmt"
	"math"
	"log"
)

type Square struct {
    side float32
}

type Circle struct {
    radius float32
}

type Shaper interface {
    Area() float32
}

func main() {
    var areaIntf Shaper
    sq1 := new(Square)
    sq1.side = 5

    areaIntf = sq1//接口绑定哪个结构体 类似虚拟类
    // Is Square the type of areaIntf?
    if t, ok := areaIntf.(*Square); ok {//如果忽略 areaIntf.(*Square) 中的 * 号，会导致编译错误：传递的是指针
		log.Println(t,ok)
        fmt.Printf("The type of areaIntf is: %T\n", t)
    }
    if u, ok := areaIntf.(*Circle); ok {
		
        fmt.Printf("The type of areaIntf is: %T\n", u)
    } else {
		log.Println(u,ok)
        fmt.Println("areaIntf does not contain a variable of type Circle")
    }
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
    return ci.radius * ci.radius * math.Pi
}

//swift模式
// switch t := areaIntf.(type) {
// case *Square:
//     fmt.Printf("Type Square %T with value %v\n", t, t)
// case *Circle:
//     fmt.Printf("Type Circle %T with value %v\n", t, t)
// case nil:
//     fmt.Printf("nil value: nothing to check?\n")
// default:
//     fmt.Printf("Unexpected type %T\n", t)
// }
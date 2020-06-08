package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    var yy *float64 = &x
    var nl []int = []int{}
    type st struct{H string}

    fmt.Println("type:", reflect.TypeOf(yy))
    fmt.Println("type:", reflect.TypeOf(nl))
    fmt.Println("type:", reflect.TypeOf(st))
    fmt.Println(len(nl))
	v := reflect.ValueOf(x)
		// setting a value:
    // v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
    fmt.Println("value:", v)
    fmt.Println("type:", v.Type())
    fmt.Println("kind:", v.Kind())
    fmt.Println("value:", v.Float())
    fmt.Println(v.Interface())
    fmt.Printf("value is %5.2e\n", v.Interface())
    y := v.Interface().(float64)
	fmt.Println(y)
	
	// v = reflect.ValueOf(&x) // Note: take the address of x.
    // fmt.Println("type of v:", v.Type())
    // fmt.Println("settability of v:", v.CanSet())
    // v = v.Elem()
    // fmt.Println("The Elem of v is: ", v)
    // fmt.Println("settability of v:", v.CanSet())
    // v.SetFloat(3.1415) // this works!
    // fmt.Println(v.Interface())
    // fmt.Println(v)
}


//反射结构体
// package main

// import (
//     "fmt"
//     "reflect"
// )

// type T struct {
//     A int
//     B string
// }

// func main() {
//     t := T{23, "skidoo"}
//     s := reflect.ValueOf(&t).Elem()
//     typeOfT := s.Type()
//     for i := 0; i < s.NumField(); i++ {
//         f := s.Field(i)
//         fmt.Printf("%d: %s %s = %v\n", i,
//             typeOfT.Field(i).Name, f.Type(), f.Interface())
//     }
//     s.Field(0).SetInt(77)
//     s.Field(1).SetString("Sunset Strip")
//     fmt.Println("t is now", t)
// }
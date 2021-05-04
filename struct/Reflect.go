package struct_

import (
	"fmt"
	"reflect"
)

func Reflect() {
	var x float64 = 3.4
	var yy *float64 = &x
	var nl []int = []int{}
	type st struct{ H string }

	fmt.Println("type:", reflect.TypeOf(yy))
	fmt.Println("type:", reflect.TypeOf(nl))
	fmt.Println("type:", reflect.TypeOf(st{"hvag"}))
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

	v = reflect.ValueOf(&x) // Note: take the address of x 传入地址 才能改变值.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem() // 通过elem方法 才能设置值
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}

//反射结构体

type Orange struct {
	size   int     `kitty:"size"`
	Weight int     `kitty:"weight"`
	Name   *string `kitty:"name"`
}

func (this Orange) GetWeight() int {
	return this.Weight
}

func ReflectStruct() {
	var name string = "hvag"
	// orange := Orange{1, 18, &name}
	orange := &Orange{1, 18, &name}

	// refValue := reflect.ValueOf(orange) // value
	// refType := reflect.TypeOf(orange)   // type => struct_.Orange
	// orangeKind := refValue.Kind()       // basic type => struct
	// fieldCount := refValue.NumField()   // field count

	// fmt.Println("orange refValue:", refValue)
	// fmt.Println("orange refType:", refType)
	// fmt.Println("orange Kind:", orangeKind)
	// fmt.Println("fieldCount:", fieldCount)

	// ptrValue := reflect.ValueOf(ptr) // value
	// ptrType := reflect.TypeOf(ptr)   // type => struct_.Orange
	// ptrValueElem := ptrValue.Elem()
	// ptrCount := ptrValueElem.NumField() // field count

	// fmt.Println("orange ptrValue:", ptrValue)
	// fmt.Println("orange ptrType:", ptrType)
	// fmt.Println("orange ptrType elem:", ptrValueElem)
	// fmt.Println("ptrCount:", ptrCount)

	// for i := 0; i < fieldCount; i++ {
	// 	fieldType := refType.Field(i)          // field type
	// 	fieldValue := refValue.Field(i)        // field vlaue
	// 	fieldTag := fieldType.Tag.Get("kitty") // field tag

	// 	fmt.Println("fieldTag:", fieldTag)
	// 	fmt.Println("field type:", fieldType.Type)
	// 	fmt.Println("fieldValue:", fieldValue)

	// }

	// // method
	// result := refValue.Method(0).Call(nil)
	// fmt.Println("method result:", result, result[0])


	// t := reflect.ValueOf(orange) // value
	// ty := reflect.TypeOf(orange)

	// switch t.Kind() {
	// case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
	// 	t = t.Elem()
	// case reflect.Struct:
	// 	for i := 0; i < t.NumField(); i++ {
	// 		field_value := t.Field(i)
	// 		field_name := ty.Field(i)
	// 		fmt.Println("name is", field_name, "value is", field_value, "and tag is", field_name.Tag)
	// 	}
	// }

	t := reflect.TypeOf(orange)
	v := reflect.ValueOf(orange)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	fieldNum := t.NumField()

    for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		kind := v.Field(i).Kind()
		value := v.Field(i)
		tag := t.Field(i).Tag
        fmt.Println("name",name,"value",value,"tag",tag,"kind",kind)
    }
	

}

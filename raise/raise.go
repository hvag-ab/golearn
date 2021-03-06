package raise

import (
	"errors"
	"fmt"
)

// error
func Calc(a int, b int) (v int, err error) {

	//捕获错误信息
	if b == 0 {
		//如果代码中出现错误 可以使用errors.New()创建错误信息
		err = errors.New("除数不能为0")
		return
	}
	v = a / b
	return

}

// panic

func Calc2(a int, b int) (v int) {

	//捕获错误信息
	if b == 0 {
		panic(errors.New("除数不能为0"))
	}
	v = a / b
	return

}

// recover
func Calc3(a int, b int) int {
	//优先使用错误拦截 在错误出现之前进行拦截 在错误出现后进行错误捕获
	//错误拦截必须配合defer使用  通过匿名函数使用
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	//捕获错误信息
	// if b == 0 {
	// 	panic(errors.New("除数不能为0"))
	// }
	v := a / b
	return v
}

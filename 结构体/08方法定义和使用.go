package main

import "fmt"

type user struct {
	name string
	email string
	age int
}


//func (方法接收者 数据类型) 方法名(方法参数列表) 返回值列表{代码体}
func (self user) notify(){//值类型 不能改变实例里面的属性值
	fmt.Println("sending user email to %s <%s>",self.name,self.email)
	self.name = "gavh"//不能改变实例self中的name属性值
}

func (self *user) change(email string){//self属性值可以改变
	self.email= email
}

func main() {

	hvag := &user{"lh","qq.com",18}
	hvag.notify()
	hvag.change("google.com")
	hvag.notify()
}

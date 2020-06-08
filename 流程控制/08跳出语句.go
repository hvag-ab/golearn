package main

import "fmt"

func main0801() {

	//死循环
	for{
		fmt.Println("跳出本层循环")

		//在循环中 跳出本层循环
		break
	}

}
func main0802(){
	//var name string
	//var passwd string
	//
	////为死循环添加break出口
	//for{
	//	fmt.Println("请输入用户名和密码")
	//	fmt.Scan(&name,&passwd)
	//
	//	if name=="admin" && passwd=="123456"{
	//		fmt.Println("成功")
	//		//跳出循环
	//		break
	//	}else{
	//		fmt.Println("输入错误")
	//	}
	//}
	sum:=0
	i:=0
	for  {
		sum+=i
		if i==100{
			break
		}
		i++
	}

	fmt.Println(sum)
}

func main0803(){

	sum:=0
	//1-100偶数的和
	for i:=1;i<=100;i++{
		if i%2==1{
			//结束本次循环 继续下次循环
			continue
		}
		sum+=i
	}

	fmt.Println(sum)
}

func main0804(){

	Loop:
	for j:=0;j<3;j++{
		fmt.Println(j)
		for a:=0;a<5;a++{
			fmt.Println(a)
			if a>3{
				break Loop
			}
		}
	}
}
//在没有使用loop标签的时候break只是跳出了第一层for循环
//使用标签后跳出到指定的标签,break只能跳出到之前，如果将Loop标签放在后边则会报错
//break标签只能用于for循环，跳出后不再执行标签对应的for循环    


func main(){
	/*
	goto语句可以无条件地转移到过程中指定的行。
通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱
goto对应(标签)既可以定义在for循环前面,也可以定义在for循环后面，当跳转到标签地方时，继续执行标签下面的代码。
	*/
   //  放在for前面，此例会一直循环下去
//    Loop://标识 变量可以自己定义
//    fmt.Println("test")
   for a:=1;a<5;a++{
	   fmt.Println(a)
	   if a>3{
		   goto Loop //goto表示跳转到哪一行运行 使用标识
	   }
   }

   Loop://放在for后边 可前可后
   fmt.Println("test")

}

func hvag(){

		i:=0
		HERE:
			print(i)
			i++
			if i==5 {
				return
			}
			goto HERE
	
}
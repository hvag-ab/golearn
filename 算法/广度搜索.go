package main

import (
	
	"fmt"


)

func nexta(value int)([]map[int]string){
	var list = []map[int]string{}
	list = append(list,map[int]string{value+2:"+2"})
	list = append(list,map[int]string{value+3:"+3"})
	list = append(list,map[int]string{value*2:"*2"})
	list = append(list,map[int]string{value*3:"*3"})

	return list
}

func solve(n int){
	var a  =[]map[int]string{}
	a =append(a,map[int]string{1:""})


	for {
		lis := a[0]
		a = a[1:]
	
		lis1 := lis
		var value int
		var path string
		for k,v :=range lis1{

			value,path =k,v
		}
		for _,p := range nexta(value){
			var (
				p1 int
				p2 string
			)
			for k1,v1 :=range p{
		
				p1,p2 =k1,v1
			}
			if p1==n{
				fmt.Println(path+p2)
				break
				return 
			}else{
				a = append(a,map[int]string{p1:path+p2})
	
			}
		}
	}

}

func main(){
	solve(75)
}

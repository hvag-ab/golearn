package main

import "fmt"

func pr(n int) []int{ 
	pt := []bool{}
	for i:=1;i<=n+1;i++{
		pt = append(pt,true)
	}
	res := []int{}
	for p:=2;p<=n;p++{
		if !pt[p]{
			continue
		}
		res = append(res,p)
		for i:=p*p;i<=n;i=i+p{
			pt[i] = false
		}
	}
	return res
}

func main(){
	fmt.Println(pr(1000000))
}

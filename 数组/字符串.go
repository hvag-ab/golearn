package main

import "fmt"
import "strings"


func main() {
    s := "\u00ff\u754c"
    for i, c := range s {
        fmt.Printf("%d:%c\n ", i, c)
	}
	b := main999()
	for _,j :=range b {
		fmt.Printf("%v",j)
	}

	stringmethod()
}

func main999() []byte{
	var b []byte
	var s string = "hvag"
	b = append(b,s...)//字符串实际是字节串数组
	return b
}

func stringmethod(){
	str:="hello"
	c:=[]byte(str)
	c[0]='c'//单引号是 字节
	s2:= string(c) // s2 == "cello"

	substr := str[0:2]
	length := len(str)

	fmt.Println(s2,substr,length)

	strs := []string{"a","b"}
	js := strings.Join(strs,"hvag")
	fmt.Print(js)



}
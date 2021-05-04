package main

import (
	"fmt"
	// "mygo/basic"
	// "mygo/file"
	// "mygo/condition"
	// "mygo/function"
	// "mygo/raise"
	"mygo/data_structure"
	// "mygo/struct"
)

func main() {
    fmt.Println("This works")
	// struct_.ReflectStruct()
	dfs := data_structure.NewDFS(0,0,8)
	fmt.Println(dfs)
	dfs.Solve()
	
}
package main

import (
	"os"
	"fmt"
	"io"
)

func main() {

	fp1, err1 := os.Open("D:\\0726浏览量\\01_Go语言（面向对象下）\\03视频\\07类型断言.avi")
	fp2, err2 := os.Create("D:/test.avi")

	if err1 != nil || err2 != nil {
		fmt.Println("文件错误")
		return
	}

	defer fp1.Close()//注意要使用defer:当打开目标文件时发生了错误，那么defer仍然能够确保src.Close()执行。如果不这么做，那么文件会一直保持打开状态并占用资源。
	defer fp2.Close()

	buf := make([]byte, 1024*1024)

	for {
		//读取的有效内容
		n, err := fp1.Read(buf)
		//将有效内容写入在文件中
		fp2.Write(buf[:n])
		if err == io.EOF {
			break
		}
	}

}

package file

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func FileBase()  {
	// 创建文件
	// fp, err := os.Create("D:/a.txt")

	f, err := os.Open("C:/itcast/test.txt")//只读方式打开
	//读写方式打开文件  1、文件路径 2、打开模式 3、打开权限 r（4）w（2）x（1）
	// fp, err := os.OpenFile("D:/a.txt", os.O_RDWR, 6) //读写追加方式打开

	if err != nil {
		fmt.Println("open err: ", err)
		return
	}
	defer f.Close()

	//如果没有出现明确的换行 writestring 会将所有的字符串写在一行
	//使用换行 需要写 \r\n windows系统中换行是以\r\n为换行 在linux中是以\n为换行
	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println("WriteString err: ", err)
		return
	}
	fmt.Println("open successful")

	// 按照行读数据
	// 获取阅读器 reader， 自带缓冲区（用户缓冲）。
	reader := bufio.NewReader(f)

	for {			// 循环读取文件， 当 err == io.EOF 结束循环
		// 使用 带分割符的 函数，读取指定数据 ‘\n’获取一行
		buf, err := reader.ReadBytes('\n')
		// 成功读取到的一行数据 保存在 buf中
		fmt.Printf("buf:%s", buf)
		if err != nil && err == io.EOF {
			break
		}
	}

}
package structgo
/**
Go语言通过首字母的大小写来控制访问权限。无论是方法，变量，常量或是自定义的变量类型，如果首字母大写，则可以被外部包访问，反之则不可以。
 */
import (
	"fmt"
)

type User struct {
	Id       int `json:"id" orm:"auto"`
	// 用户名
	Username string `json:"username"`
}

func main0() {
	//值类型
	u1:=models.User{}
	var u2  models.User

	//指针类型
	u3:=new(models.User)
	u4:=&models.User{}
	var u5  *models.User

	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)

	fmt.Println(u4)
	fmt.Println(u5)
}
/*
输出：

{0  }
 {0  }
 &{0  }
 &{0  }
 <nil>
*/
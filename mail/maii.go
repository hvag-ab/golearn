package main

import (
    "gopkg.in/gomail.v2"
	"strconv"
	"fmt"
)

func SendMail(mailTo []string,subject string, body string ) error {
  //定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
    mailConn := map[string]string {
        "user": "39435773qq@sina.cn", 
        "pass": "bdc088b76c9f9230",  
        "host": "smtp.sina.cn",
        "port": "465",
    }

    port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

    m := gomail.NewMessage()
    m.SetHeader("From","hvag" + "<" + mailConn["user"] + ">")  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Cc","475886877@qq.com","475886877@qq.com") //抄送 多个邮件地址
    // m.SetHeader("Bcc","475886877@qq.com","lh") //暗送
    m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文 // 可以放html..还有其他的
	// m.SetBody("text/plain","我是正文") // 正文
	// m.Attach("/Applications/HvagPicture/HVAG/myself/1.jpg")  //添加附件
	// m.Attach("/Applications/HvagPicture/HVAG/myself/何雪.JPG")  //添加附件

    d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

    err := d.DialAndSend(m)
    return err

}
func main()  {
   //定义收件人
     mailTo := []string {
    // "39435773qq@sina.cn",
    "1603753920@qq.com",
    }
   //邮件主题为"Hello"
    subject := "Hello"
   // 邮件正文
    body := "Good"
	err:= SendMail(mailTo, subject, body)
	fmt.Println(err)
}
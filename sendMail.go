package main



import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main(){
	//创建新的发送消息
	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	//设置发送人
	m.SetAddressHeader("From", "@pactera.com", "winnie")
	//设置接收人
	var to []string
	to = append(to, m.FormatAddress("@qq.com", "huxiaoyu"))
	m.SetHeader("To", to...)
	//设置邮件名
	m.SetHeader("subject", "测试")
	//设置邮件内容
	m.SetBody("text/Plain", "测试")
	//发送
	d := gomail.NewDialer("outlook.office365.com", 587, "@pactera.com", "")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify:true}
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
}
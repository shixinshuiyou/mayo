package mail

import (
	"log"

	"github.com/go-gomail/gomail"
)

type EmailInfo struct {
	ServerHost string // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerPort int    // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465

	FromEmail  string // FromEmail　发件人邮箱地址
	FromPasswd string //发件人邮箱密码（注意，这里是明文形式)

	Recipient []string //收件人邮箱
	CC        []string //抄送
}

/**
 * @Author: czh
 * @Date: 2021-07-06 15:45:55
 * @Description: 发送邮件
 * @Param : subject[主题]、body[内容]、emailInfo[发邮箱需要的信息(参考EmailInfo)]
 * @Return:
 */
func SendEmail(subject, body string, emailInfo *EmailInfo) {
	if len(emailInfo.Recipient) == 0 {
		log.Print("收件人列表为空")
		return
	}

	mes := gomail.NewMessage()
	//设置收件人
	mes.SetHeader("To", emailInfo.Recipient...)
	//设置抄送列表
	if len(emailInfo.CC) != 0 {
		mes.SetHeader("Cc", emailInfo.CC...)
	}
	// 第三个参数为发件人别名，如"dcj"，可以为空（此时则为邮箱名称）
	mes.SetAddressHeader("From", emailInfo.FromEmail, "dcj")

	//主题
	mes.SetHeader("Subject", subject)

	//正文
	mes.SetBody("text/html", body)

	d := gomail.NewDialer(emailInfo.ServerHost, emailInfo.ServerPort, emailInfo.FromEmail, emailInfo.FromPasswd)
	err := d.DialAndSend(mes)
	if err != nil {
		log.Println("发送邮件失败： ", err)
	} else {
		log.Println("已成功发送邮件到指定邮箱")
	}
}

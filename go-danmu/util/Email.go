package util

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"

	"github.com/spf13/viper"
)

const (
	RegisterCode = "注册验证"
)

//参数: 目标邮箱，验证码，主题
func SendEmail(ToEmail string, code string, subject string) bool {
	host := viper.GetString("email.host")
	port := viper.GetInt("email.port")
	email := viper.GetString("email.address")
	password := viper.GetString("email.password")
	toEmail := ToEmail
	header := make(map[string]string)
	header["From"] = viper.GetString("email.name") + "<" + email + ">"
	header["To"] = toEmail
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := "<h3>尊敬的用户：</h3><p>您好! 您的验证码是 <span style='color:red'> " + code + "</span>，五分钟内有效，祝您生活愉快！</p>"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		email,
		[]string{toEmail},
		[]byte(message),
	)
	if err != nil {
		return false
	}
	return true
}

func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	c, err := dial(addr)
	if err != nil {
		LogError("Create smpt client error:" + err.Error())
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				LogError("Error during AUTH:" + err.Error())
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		LogError("Dialing Error:" + err.Error())
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

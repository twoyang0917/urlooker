package tests

import (
	"testing"

	"github.com/toolkits/smtp"
)

const (
	tos     = "xxxxxxxx@xxxx.com"
	subject = "test"
	body    = "test"
)

func Test_SendMail(t *testing.T) {
	address := "smtp.exmail.qq.com:465"
	from := "xxxxxxxx@xxxx.com"
	password := "xxxxxxxx"
	tls := true
	anonymous := false
	skipVerify := false
	s := smtp.NewSMTP(address, from, password, tls, anonymous, skipVerify)
	t.Log(s.SendMail(from, tos, subject, body))
}

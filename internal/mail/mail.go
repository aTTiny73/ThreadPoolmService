package mail

import (
	"encoding/json"
	"fmt"

	"github.com/aTTiny73/ThreadPoolmService/pkg/pool"
	"github.com/irnes/go-mailer"
)

// Mailer function sends an email to given recipients
func Mailer(mailData []byte) {

	data := pool.Host{}
	err := json.Unmarshal(mailData, &data)
	if err != nil {
		fmt.Println(err)
	}

	var (
		host = "xxx"
		user = "xxx"
		pass = "xxx"
	)
	config := mailer.Config{
		Host: host,
		Port: 465,
		User: user,
		Pass: pass,
	}
	Mailer := mailer.NewMailer(config, true)

	mail := mailer.NewMail()
	mail.FromName = "Go Mailer"
	mail.From = user

	for _, mailAddres := range data.Recipients {

		mail.SetTo(mailAddres)
		mail.Subject = "Server "
		mail.Body = "Your server is down"

		if err := Mailer.Send(mail); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Mail successfuly sent to", mailAddres)
		}
	}

}

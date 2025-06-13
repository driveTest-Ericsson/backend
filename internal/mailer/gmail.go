package mailer

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

type GmailMailer struct {
	fromEmail string
	password  string
}

func NewGmail(fromEmail, password string) *GmailMailer {

	return &GmailMailer{
		fromEmail: fromEmail,
		password:  password,
	}
}

type GmailData struct {
	Username      string
	ActivationURL string
}

func (m *GmailMailer) Send(recipient string, data *GmailData) error {

	from := m.fromEmail
	pass := m.password

	to := recipient

	// // template parsing and building
	// tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	// if err != nil {
	// 	return err
	// }

	// subject := new(bytes.Buffer)
	// err = tmpl.ExecuteTemplate(subject, "subject", data)
	// if err != nil {
	// 	return err
	// }

	// body := new(bytes.Buffer)
	// err = tmpl.ExecuteTemplate(body, "body", data)
	// if err != nil {
	// 	return err
	// }

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Welcome to My Website")
	message.SetBody("text/html",
		`
		<!doctype html>
		<html>
			<head>
				<meta name="viewport" content="width=device-width" />
				<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			</head>
			<body>
				<p>Hi `+data.Username+`,</p>
				<p>Thanks for signing up for ArminTest, excited to have you on board!</p>
				<p>
					Before you can start using ArminTest, you need to confirm your email address. Click the link below
					to confirm your email address:
				</p>
				<p><a href="`+data.ActivationURL+`">`+data.ActivationURL+`</a></p>
				<p>If you didn't sign up for ArminTest, you can safely ignore this email.</p>

				<p>Thanks,</p>
				<p>ArminTest</p>
			</body>
		</html>
		`,
	)
	message.AddAlternative("text/plain", "Hi "+data.Username+",\nPlease confirm your email: "+data.ActivationURL)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

	var retryError error
	for i := 0; i < maxRetries; i++ {
		retryError = d.DialAndSend(message)
		if retryError != nil {
			// exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		return nil
	}

	return fmt.Errorf("failed to send email after %d attempts, error: %v", maxRetries, retryError)
}

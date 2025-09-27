package sendgrid

import (
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridProvider struct {
	ApiKey        string
	SenderEmail   string
	ReceiverEmail string
}

func (s SendGridProvider) SendEmail(newIp string) error {
	if s.ApiKey == "" {
		println("SendGrid not configured, skipping email...")
		return nil
	}

	from := mail.NewEmail("IP Checker", s.SenderEmail)
	subject := "New IP Address"
	to := mail.NewEmail("Friend", s.ReceiverEmail)

	plainTextContent := "Your new IP address is: " + newIp
	htmlContent := "<strong>Your new IP address is: " + newIp + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(s.ApiKey)
	response, err := client.Send(message)
	if err != nil || response.StatusCode != 202 {
		println("error sending email. Status code " + strconv.Itoa((response.StatusCode)))
		println(err.Error())
		return err
	}
	println("email sent successfully")
	return nil
}

func (s SendGridProvider) SendErrorEmail() error {
	if s.ApiKey == "" {
		println("SendGrid not configured, skipping email...")
		return nil
	}

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."

	from := mail.NewEmail("IP Checker", s.SenderEmail)
	subject := "IP Checker Error"
	to := mail.NewEmail("Friend", s.ReceiverEmail)

	plainTextContent := errMess
	htmlContent := "<strong>" + errMess + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(s.ApiKey)
	response, err := client.Send(message)
	if err != nil {
		println(err)
		return err
	} else {
		println(response.StatusCode)
	}
	return nil
}

func (s SendGridProvider) SendDnsErrorEmail() error {
	errMess := "There was an error updating your DNS record. Please check your internet connection and try again."

	from := mail.NewEmail("IP Checker", s.SenderEmail)
	subject := "IP Checker Error"
	to := mail.NewEmail("Friend", s.ReceiverEmail)

	plainTextContent := errMess
	htmlContent := "<strong>" + errMess + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(s.ApiKey)
	response, err := client.Send(message)
	if err != nil {
		println(err)
		return err
	} else {
		println(response.StatusCode)
	}
	return nil
}

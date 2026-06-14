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

func (s SendGridProvider) send(subject, text, html string) error {
	from := mail.NewEmail("IP Checker", s.SenderEmail)
	to := mail.NewEmail("Friend", s.ReceiverEmail)

	message := mail.NewSingleEmail(from, subject, to, text, html)

	client := sendgrid.NewSendClient(s.ApiKey)
	response, err := client.Send(message)
	if err != nil || response.StatusCode != 202 {
		println("error sending email. Status code " + strconv.Itoa(response.StatusCode))
		return err
	}
	return nil
}

func (s SendGridProvider) SendEmail(newIp string) error {
	if s.ApiKey == "" {
		println("SendGrid not configured, skipping email...")
		return nil
	}

	if err := s.send("New IP Address", "Your new IP address is: "+newIp, "<strong>Your new IP address is: "+newIp+"</strong>"); err != nil {
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
	if err := s.send("IP Checker Error", errMess, "<strong>"+errMess+"</strong>"); err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func (s SendGridProvider) SendDnsErrorEmail() error {
	errMess := "There was an error updating your DNS record. Please check your internet connection and try again."
	if err := s.send("IP Checker Error", errMess, "<strong>"+errMess+"</strong>"); err != nil {
		println(err.Error())
		return err
	}
	return nil
}

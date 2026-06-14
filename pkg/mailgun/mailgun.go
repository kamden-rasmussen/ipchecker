package mailgun

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunProvider struct {
	ApiKey        string
	Domain        string
	SenderEmail   string
	ReceiverEmail string
}

func (m MailgunProvider) send(subject, text, html string) error {
	mg := mailgun.NewMailgun(m.Domain, m.ApiKey)

	message := mg.NewMessage(m.SenderEmail, subject, text, m.ReceiverEmail)
	message.SetHtml(html)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		println("error sending email: " + err.Error())
		return err
	}

	println("email sent successfully. ID: " + id + ", Response: " + resp)
	return nil
}

func (m MailgunProvider) SendEmail(newIp string) error {
	if m.ApiKey == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	return m.send("New IP Address", "Your new IP address is: "+newIp, "<strong>Your new IP address is: "+newIp+"</strong>")
}

func (m MailgunProvider) SendErrorEmail() error {
	if m.ApiKey == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."
	return m.send("IP Checker Error", errMess, "<strong>"+errMess+"</strong>")
}

func (m MailgunProvider) SendDnsErrorEmail() error {
	errMess := "There was an error updating your DNS record. Please check your internet connection and try again."
	return m.send("IP Checker Error", errMess, "<strong>"+errMess+"</strong>")
}

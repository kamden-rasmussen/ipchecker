package mailgun

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunProvider struct {
	APIKey        string
	Domain        string
	SenderEmail   string
	ReceiverEmail string
}

func (m MailgunProvider) send(subject, text, html string) error {
	mg := mailgun.NewMailgun(m.Domain, m.APIKey)

	message := mailgun.NewMessage(m.SenderEmail, subject, text, m.ReceiverEmail)
	message.SetHTML(html)

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

func (m MailgunProvider) SendEmail(newIP string) error {
	if m.APIKey == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	return m.send("New IP Address", "Your new IP address is: "+newIP, "<strong>Your new IP address is: "+newIP+"</strong>")
}

func (m MailgunProvider) SendErrorEmail() error {
	if m.APIKey == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."
	return m.send("IP Checker Error", errMess, "<strong>"+errMess+"</strong>")
}

func (m MailgunProvider) SendDNSErrorEmail() error {
	errMess := "There was an error updating your DNS record. Please check your internet connection and try again."
	return m.send("IP Checker Error", errMess, "<strong>"+errMess+"</strong>")
}

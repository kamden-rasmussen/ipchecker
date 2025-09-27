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

func (m MailgunProvider) SendEmail(newIp string) error {
	if m.ApiKey == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	mg := mailgun.NewMailgun(m.Domain, m.ApiKey)

	message := mg.NewMessage(m.SenderEmail, "New IP Address", "Your new IP address is: "+newIp, m.ReceiverEmail)
	message.SetHtml("<strong>Your new IP address is: " + newIp + "</strong>")

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

func (m MailgunProvider) SendErrorEmail() error {
	if m.ApiKey == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	mg := mailgun.NewMailgun(m.Domain, m.ApiKey)

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."

	message := mg.NewMessage(m.SenderEmail, "IP Checker Error", errMess, m.ReceiverEmail)
	message.SetHtml("<strong>" + errMess + "</strong>")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		println("error sending error email: " + err.Error())
		return err
	}

	println("error email sent successfully. ID: " + id + ", Response: " + resp)
	return nil
}

func (m MailgunProvider) SendCloudflareErrorEmail() error {
	mg := mailgun.NewMailgun(m.Domain, m.ApiKey)

	errMess := "There was an error updating your Cloudflare DNS record. Please check your internet connection and try again."

	message := mg.NewMessage(m.SenderEmail, "IP Checker Error", errMess, m.ReceiverEmail)
	message.SetHtml("<strong>" + errMess + "</strong>")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		println("error sending cloudflare error email: " + err.Error())
		return err
	}

	println("cloudflare error email sent successfully. ID: " + id + ", Response: " + resp)
	return nil
}

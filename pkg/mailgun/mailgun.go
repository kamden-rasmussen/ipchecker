package mailgun

import (
	"context"
	"time"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
	"github.com/mailgun/mailgun-go/v4"
)

type MailgunProvider struct{}

func (m MailgunProvider) SendEmail(newIp string) error {
	if env.GetKey("MAILGUN_API_KEY") == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	mg := mailgun.NewMailgun(env.GetKey("MAILGUN_DOMAIN"), env.GetKey("MAILGUN_API_KEY"))

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	message := mg.NewMessage(senderEmail, "New IP Address", "Your new IP address is: "+newIp, receiverEmail)
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
	if env.GetKey("MAILGUN_API_KEY") == "" {
		println("Mailgun not configured, skipping email...")
		return nil
	}

	mg := mailgun.NewMailgun(env.GetKey("MAILGUN_DOMAIN"), env.GetKey("MAILGUN_API_KEY"))

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	message := mg.NewMessage(senderEmail, "IP Checker Error", errMess, receiverEmail)
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
	mg := mailgun.NewMailgun(env.GetKey("MAILGUN_DOMAIN"), env.GetKey("MAILGUN_API_KEY"))

	errMess := "There was an error updating your Cloudflare DNS record. Please check your internet connection and try again."

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	message := mg.NewMessage(senderEmail, "IP Checker Error", errMess, receiverEmail)
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

func SendEmail(newIp string) error {
	provider := MailgunProvider{}
	return provider.SendEmail(newIp)
}

func SendErrorEmail() error {
	provider := MailgunProvider{}
	return provider.SendErrorEmail()
}

func SendCloudflareErrorEmail() error {
	provider := MailgunProvider{}
	return provider.SendCloudflareErrorEmail()
}

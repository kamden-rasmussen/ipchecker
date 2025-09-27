package sendgrid

import (
	"strconv"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridProvider struct{}

func (s SendGridProvider) SendEmail(newIp string) error {
	if env.GetKey("SENDGRID_API_KEY") == "" {
		println("SendGrid not configured, skipping email...")
		return nil
	}

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "New IP Address"
	to := mail.NewEmail("Friend", receiverEmail)

	plainTextContent := "Your new IP address is: " + newIp
	htmlContent := "<strong>Your new IP address is: " + newIp + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(env.GetKey("SENDGRID_API_KEY"))
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
	if env.GetKey("SENDGRID_API_KEY") == "" {
		println("SendGrid not configured, skipping email...")
		return nil
	}

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "IP Checker Error"
	to := mail.NewEmail("Friend", receiverEmail)

	plainTextContent := errMess
	htmlContent := "<strong>" + errMess + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(env.GetKey("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		println(err)
		return err
	} else {
		println(response.StatusCode)
	}
	return nil
}

func (s SendGridProvider) SendCloudflareErrorEmail() error {
	errMess := "There was an error updating your Cloudflare DNS record. Please check your internet connection and try again."

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "IP Checker Error"
	to := mail.NewEmail("Friend", receiverEmail)

	plainTextContent := errMess
	htmlContent := "<strong>" + errMess + "</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(env.GetKey("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		println(err)
		return err
	} else {
		println(response.StatusCode)
	}
	return nil
}

func SendEmail(newIp string) error {
	provider := SendGridProvider{}
	return provider.SendEmail(newIp)
}

func SendErrorEmail() error {
	provider := SendGridProvider{}
	return provider.SendErrorEmail()
}

func SendCloudflareErrorEmail() error {
	provider := SendGridProvider{}
	return provider.SendCloudflareErrorEmail()
}

package email

import (
	"strconv"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(newIp string) error {
	if env.GetKey("SENDGRID_API_KEY") == "" {
		// Don't send an email if they don't have it defined
		return nil
	}

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	// email
	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "New IP Address"
	to := mail.NewEmail("Friend", receiverEmail)

	// body
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

func SendErrorEmail() error {
	if env.GetKey("SENDGRID_API_KEY") == "" {
		// Don't send an email if they don't have it defined
		return nil
	}

	errMess := "There was an error checking your IP address. Please check your internet connection and try again."

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	// email
	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "IP Checker Error"
	to := mail.NewEmail("Friend", receiverEmail)

	// body
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

func SendCloudflareErrorEmail() error {

	errMess := "There was an error updating your Cloudflare DNS record. Please check your internet connection and try again."

	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	// email
	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "IP Checker Error"
	to := mail.NewEmail("Friend", receiverEmail)

	// body
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

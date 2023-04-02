package email

import (
	"log"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(newIp string) error {
	
	senderEmail := env.GetKey("SENDER_EMAIL")
	receiverEmail := env.GetKey("RECEIVER_EMAIL")

	// email
	from := mail.NewEmail("IP Checker", senderEmail)
	subject := "New IP Address"
	to := mail.NewEmail("Kamden", receiverEmail)
	
	// body
	plainTextContent := "Your new IP address is: " + newIp
	htmlContent := "<strong>Your new IP address is: " + newIp + "</strong>"
	
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	
	client := sendgrid.NewSendClient(env.GetKey("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
	return nil
}
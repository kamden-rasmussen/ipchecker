package resend

import (
	resendgo "github.com/resend/resend-go/v2"
)

type ResendProvider struct {
	ApiKey        string
	SenderEmail   string
	ReceiverEmail string
}

func (r ResendProvider) send(subject, text string) error {
	client := resendgo.NewClient(r.ApiKey)

	params := &resendgo.SendEmailRequest{
		From:    r.SenderEmail,
		To:      []string{r.ReceiverEmail},
		Subject: subject,
		Html:    "<strong>" + text + "</strong>",
		Text:    text,
	}

	_, err := client.Emails.Send(params)
	return err
}

func (r ResendProvider) SendEmail(newIp string) error {
	if r.ApiKey == "" {
		println("Resend not configured, skipping email...")
		return nil
	}

	if err := r.send("New IP Address", "Your new IP address is: "+newIp); err != nil {
		println("error sending email: " + err.Error())
		return err
	}

	println("email sent successfully")
	return nil
}

func (r ResendProvider) SendErrorEmail() error {
	if r.ApiKey == "" {
		println("Resend not configured, skipping email...")
		return nil
	}

	if err := r.send("IP Checker Error", "There was an error checking your IP address. Please check your internet connection and try again."); err != nil {
		println("error sending error email: " + err.Error())
		return err
	}

	println("error email sent successfully")
	return nil
}

func (r ResendProvider) SendDnsErrorEmail() error {
	if err := r.send("IP Checker Error", "There was an error updating your DNS record. Please check your internet connection and try again."); err != nil {
		println("error sending DNS error email: " + err.Error())
		return err
	}

	println("DNS error email sent successfully")
	return nil
}

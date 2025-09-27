package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kamden-rasmussen/ipchecker/pkg/check"
	"github.com/kamden-rasmussen/ipchecker/pkg/cloudflare"
	"github.com/kamden-rasmussen/ipchecker/pkg/cron"
	"github.com/kamden-rasmussen/ipchecker/pkg/env"
	"github.com/kamden-rasmussen/ipchecker/pkg/godaddy"
	"github.com/kamden-rasmussen/ipchecker/pkg/mailgun"
	"github.com/kamden-rasmussen/ipchecker/pkg/sendgrid"
)

type DnsHost interface {
	PutNewIP(string) (int, error)
}

type EmailProvider interface {
	SendEmail(string) error
	SendErrorEmail() error
	SendDnsErrorEmail() error
}

func main() {
	// open log file
	openLogFile()
	println("Starting IP Checker")

	// load env
	env.InitEnv()

	RunCheck()

	// set up cron jobs
	cronService := cron.NewCron()
	cronService.AddFunc("*/5 * * * *", RunCheck)
	cronService.Start()

	// run forever
	forever()
}

func openLogFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func RunCheck() {
	var dnsHost DnsHost
	var emailProvider EmailProvider

	dnsProvider := env.GetKey("DNSHOST")
	emailProviderType := env.GetKey("EMAIL_PROVIDER")

	// Get bool env vars
	shouldUpdate, err := strconv.ParseBool(env.GetKey("UPDATEDNS"))
	if err != nil {
		println(err.Error())
		return
	}

	// Set up DNS provider
	switch strings.ToUpper(dnsProvider) {
	case "CLOUDFLARE":
		dnsHost = cloudflare.Cloudflare{
			env.GetKey("CLOUDFLARE_ZONE_ID"),
			env.GetKey("CLOUDFLARE_DNS_ID"),
			env.GetKey("CLOUDFLARE_EMAIL"),
			"Bearer " + env.GetKey("CLOUDFLARE_API_KEY"),
			env.GetKey("CLOUDFLARE_DOMAIN_NAME"),
		}
	case "GODADDY":
		dnsHost = godaddy.Godaddy{
			env.GetKey("GODADDY_DOMAIN_NAME"),
			env.GetKey("GODADDY_DNS_RECORD_TYPE"),
			env.GetKey("GODADDY_DNS_RECORD_NAME"),
			env.GetKey("GODADDY_API_KEY"),
			env.GetKey("GODADDY_API_SECRET"),
		}
	case "NA": // if you do not have a DNS set up and just want the email
		break
	default:
		println("DNS provider not implemented yet")
		return
	}

	// Set up email provider
	switch strings.ToUpper(emailProviderType) {
	case "SENDGRID":
		emailProvider = sendgrid.SendGridProvider{
			ApiKey:        env.GetKey("SENDGRID_API_KEY"),
			SenderEmail:   env.GetKey("SENDER_EMAIL"),
			ReceiverEmail: env.GetKey("RECEIVER_EMAIL"),
		}
	case "MAILGUN":
		emailProvider = mailgun.MailgunProvider{
			ApiKey:        env.GetKey("MAILGUN_API_KEY"),
			Domain:        env.GetKey("MAILGUN_DOMAIN"),
			SenderEmail:   env.GetKey("SENDER_EMAIL"),
			ReceiverEmail: env.GetKey("RECEIVER_EMAIL"),
		}
	default:
		println("Email provider not implemented yet")
		return
	}

	ip := check.CheckIp()
	if ip == "outage" {
		err := emailProvider.SendErrorEmail()
		if err != nil {
			println(err.Error())
		}
		return
	}
	if ip != "" {
		err := emailProvider.SendEmail(ip)
		if err != nil {
			println(err.Error())
		}

		if shouldUpdate {
			code, err := dnsHost.PutNewIP(ip)
			if err != nil || code != 200 {
				fmt.Printf("Failed to update DNS record. Status code %d\n", code)
				emailProvider.SendDnsErrorEmail()
			} else {
				println("Successfully updated DNS record")
			}
		}
	}
}

func forever() {
	select {}
}

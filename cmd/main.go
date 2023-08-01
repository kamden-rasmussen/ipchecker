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
	"github.com/kamden-rasmussen/ipchecker/pkg/email"
	"github.com/kamden-rasmussen/ipchecker/pkg/env"
	"github.com/kamden-rasmussen/ipchecker/pkg/godaddy"
)

type DnsHost interface {
	PutNewIP(string) (int, error)
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
	dnsProvider := env.GetKey("DNSHOST")

	// Get bool env vars
	shouldUpdate, err := strconv.ParseBool(env.GetKey("UPDATEDNS"))
	if err != nil {
		println(err.Error())
		return
	}

	switch strings.ToUpper(dnsProvider) {
	case "CLOUDFLARE":
		dnsHost = cloudflare.Cloudflare{
			env.GetKey("ZONE_ID"),
			env.GetKey("DNS_ID"),
			env.GetKey("EMAIL"),
			"Bearer " + env.GetKey("API_KEY"),
			env.GetKey("DOMAIN_NAME"),
		}
	case "GODADDY":
		dnsHost = godaddy.Godaddy{
			env.GetKey("DOMAIN_NAME"),
			env.GetKey("DNS_RECORD_TYPE"),
			env.GetKey("DNS_RECORD_NAME"),
			env.GetKey("API_KEY"),
			env.GetKey("API_SECRET"),
		}
	default:
		println("Not implemented yet")
		return
	}

	ip := check.CheckIp()
	if ip == "outage" {
		err := email.SendErrorEmail()
		if err != nil {
			println(err.Error())
		}
		return
	}
	if ip != "" {
		err := email.SendEmail(ip)
		if err != nil {
			println(err.Error())
		}

		if shouldUpdate {
			code, err := dnsHost.PutNewIP(ip)
			if err != nil || code != 200 {
				fmt.Printf("Failed to update DNS record. Status code %d\n", code)
				email.SendCloudflareErrorEmail()
			} else {
				println("Successfully updated DNS record")
			}
		}
	}
}

func forever() {
	select {}
}

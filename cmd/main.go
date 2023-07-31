package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kamden-rasmussen/ipchecker/pkg/check"
	"github.com/kamden-rasmussen/ipchecker/pkg/cloudflare"
	"github.com/kamden-rasmussen/ipchecker/pkg/cron"
	"github.com/kamden-rasmussen/ipchecker/pkg/email"
	"github.com/kamden-rasmussen/ipchecker/pkg/env"
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
	dnsProvider := os.Getenv("DNSHOST")

	// turn UPDATEDNS env var into a bool
	shouldUpdate, err := strconv.ParseBool(os.Getenv("UPDATEDNS"))
	if err != nil {
		println("Bool err: ", err.Error())
		return
	}

	switch dnsProvider {
	case "Cloudflare":
		dnsHost = cloudflare.Cloudflare{
			env.GetKey("ZONE_ID"),
			env.GetKey("DNS_ID"),
			env.GetKey("EMAIL"),
			"Bearer " + env.GetKey("API_KEY"),
			env.GetKey("DOMAIN_NAME"),
		}
	default:
		fmt.Println("Not implemented yet")
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
				println("Failed to update Cloudflare DNS record. Status code " + strconv.Itoa(code))
				email.SendCloudflareErrorEmail()
			} else {
				println("Successfully updated Cloudflare DNS record")
			}
		}
	}
}

func forever() {
	select {}
}

package main

import (
	"log"
	"os"
	"strconv"

	"github.com/kamden-rasmussen/ipchecker/pkg/check"
	"github.com/kamden-rasmussen/ipchecker/pkg/cloudflare"
	"github.com/kamden-rasmussen/ipchecker/pkg/cron"
	"github.com/kamden-rasmussen/ipchecker/pkg/email"
	"github.com/kamden-rasmussen/ipchecker/pkg/env"
)

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
	shouldUseCloudflare := os.Getenv("CLOUDFLARE")
	boolCloudflare := false
	var err error
	// turn shouldUseCloudflare into a bool
	if shouldUseCloudflare != "" {
		boolCloudflare, err = strconv.ParseBool(shouldUseCloudflare)
		if err != nil {
			println(err)
		}
	}

	ip := check.CheckIp()
	if ip == "outage" {
		err := email.SendErrorEmail()
		if err != nil {
			println(err)
		}
		return
	}
	if ip != "" {
		err := email.SendEmail(ip)
		if err != nil {
			println(err)
		}
		if boolCloudflare {
			code, err := cloudflare.PutNewIP(ip)
			if err != nil || code != 200 {
				println("failed to update Cloudflare DNS record. Status code " + strconv.Itoa(code))
				email.SendCloudflareErrorEmail()
			} else {
				println("successfully updated Cloudflare DNS record")
			}
		}
	}
}

func forever() {
	select {}
}

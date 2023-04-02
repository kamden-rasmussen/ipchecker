package main

import (
	"log"
	"os"

	"github.com/kamden-rasmussen/ipchecker/pkg/check"
	"github.com/kamden-rasmussen/ipchecker/pkg/cron"
	"github.com/kamden-rasmussen/ipchecker/pkg/email"
	"github.com/kamden-rasmussen/ipchecker/pkg/env"
)

func main() {
	// open log file
	openLogFile()

	// load env
	env.InitEnv()

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

func RunCheck(){
	ip := check.CheckIp()
	if ip != "" {
		err := email.SendEmail(ip)
		if err != nil {
			log.Println(err)
		}
	}
}

func forever() {
	select {}
}

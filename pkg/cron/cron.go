package cron

import (
	"log"

	"github.com/robfig/cron/v3"
)

type Cron struct {
	cronService *cron.Cron
}

func NewCron() *Cron {
	cronService := cron.New()
	return &Cron{cronService: cronService}
}

func (c *Cron) AddFunc(spec string, cmd func()) {
	_, err := c.cronService.AddFunc(spec, cmd)
	if err != nil {
		log.Printf("error adding cron service func: %v", err)
	}
}

func (c *Cron) Start() {
	c.cronService.Start()
}

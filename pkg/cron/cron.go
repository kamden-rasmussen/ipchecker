package cron

import "github.com/robfig/cron/v3"

type Cron struct {
	cronService *cron.Cron
}

func NewCron() *Cron{
	cronService := cron.New()
	return &Cron{cronService: cronService}
}

func (c *Cron) AddFunc(spec string, cmd func()) {
	c.cronService.AddFunc(spec, cmd)
}

func (c *Cron) Start() {
	c.cronService.Start()
}
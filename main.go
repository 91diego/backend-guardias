package main

import (
	"fmt"

	"github.com/91diego/backend-guardias/src/routes"
	"github.com/91diego/backend-guardias/src/utils"
	"github.com/robfig/cron/v3"
)

func main() {

	utils.EnvVariables()

	//log.Info("Create new cronjob")
	//c := cron.New()
	//c.AddFunc("*/1 * * * *", func() { log.Info("[Job 1]Every minute job\n") })

	// start cron with one scheduled job
	//log.Info("Start cron")
	//utils.PrintCronEntries(c.Entries())
	//time.Sleep(2 * time.Minute)

	// Funcs may also be added to a running Cron
	//log.Info("Add new job to a running cron")
	//entryID2, _ := c.AddFunc("*/2 * * * *", func() { log.Info("[Job 2]Every two minutes job\n") })
	//utils.PrintCronEntries(c.Entries())
	//time.Sleep(5 * time.Minute)

	//Remove Job2 and add new Job2 that run every 1 minute
	//log.Info("Remove Job2 and add new Job2 with schedule run every minute")
	//c.Remove(entryID2)
	//c.AddFunc("*/1 * * * *", func() { log.Info("[Job 2]Every one minute job\n") })
	//time.Sleep(5 * time.Minute)
	fmt.Println("JOB CRON")
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { fmt.Println("[Job 1]Every minute job\n") })
	// start cron with one scheduled job
	fmt.Println("Start cron")
	utils.PrintCronEntries(c.Entries())
	routes.Router()
}

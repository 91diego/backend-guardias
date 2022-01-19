package utils

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func PrintCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}

package main

import (
	"fmt"
	"time"

	"github.com/91diego/backend-guardias/src/models"
	"github.com/91diego/backend-guardias/src/repositories"
	"github.com/91diego/backend-guardias/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {

	// utils.EnvVariables()
	mexicoCity, _ := time.LoadLocation("America/Mexico_City")
	c := cron.New(cron.WithLocation(mexicoCity))
	c.AddFunc("00 10 * * *", CheckGuards)
	c.AddFunc("30 14 * * *", CheckGuards)
	c.AddFunc("12 20 * * *", CheckGuards)
	c.Start()

	gin.SetMode(gin.ReleaseMode)
	routes.Router()
}

func CheckGuards() {

	var advisoryGuardByShift []models.AdvisorGuard
	var advisoryGuards []models.AdvisorGuard
	layout := "2006-01-02 15:04:05"
	now := time.Now().UTC().Local()

	currentDate := fmt.Sprintf("%d-%02d-%02d",
		now.Year(), now.Month(), now.Day())

	morningChangeLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	afternoonChangeLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	eveningChangeLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	morningGuardChange := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"10", "00", "00")
	afternoonGuardChange := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"14", "30", "00")
	eveningGuardChange := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"20", "12", "00")

	morningLayoutParse, err := time.Parse(layout, morningChangeLayout)
	if err != nil {
		return
	}
	afternoonLayoutParse, err := time.Parse(layout, afternoonChangeLayout)
	if err != nil {
		return
	}

	eveningLayoutParse, err := time.Parse(layout, eveningChangeLayout)
	if err != nil {
		return
	}

	morningParse, err := time.Parse(layout, morningGuardChange)
	if err != nil {
		return
	}
	afternoonParse, err := time.Parse(layout, afternoonGuardChange)
	if err != nil {
		return
	}

	eveningParse, err := time.Parse(layout, eveningGuardChange)
	if err != nil {
		return
	}

	if morningParse == morningLayoutParse {
		repositories.New().GetAdvisoryGuardByShift("MATUTINO", currentDate, &advisoryGuardByShift)
	}

	if afternoonParse == afternoonLayoutParse {
		repositories.New().GetAdvisoryGuardByShift("VESPERTINO", currentDate, &advisoryGuardByShift)
	}

	if eveningParse == eveningLayoutParse {
		repositories.New().GetAdvisoryGuardByShift("NOCTURNO", currentDate, &advisoryGuardByShift)
	}

	// Add flag on B24
	for _, v := range advisoryGuardByShift {
		advisorBitrix := models.AdvisorBitrix{
			UserID:        v.AdvisorBitrixID,
			PersonalSreet: "GUARDIA " + v.Development,
		}
		_ = models.UpdateBitrixGuardAdvisor(&advisorBitrix)
	}

	// Delete flag on B24
	repositories.New().GetAdvisoryGuardsDB(currentDate, &advisoryGuards)
	for _, v := range advisoryGuards {
		advisorBitrix := models.AdvisorBitrix{
			UserID:        v.AdvisorBitrixID,
			PersonalSreet: "",
		}
		_ = models.UpdateBitrixGuardAdvisor(&advisorBitrix)
	}
}

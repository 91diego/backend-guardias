package main

import (
	"fmt"
	"time"

	"github.com/91diego/backend-guardias/src/models"
	"github.com/91diego/backend-guardias/src/repositories"
	"github.com/91diego/backend-guardias/src/routes"
	"github.com/robfig/cron/v3"
)

func main() {

	// utils.EnvVariables()
	mexicoCity, _ := time.LoadLocation("America/Mexico_City")
	c := cron.New(cron.WithLocation(mexicoCity))
	c.AddFunc("00 10 * * *", CheckGuards)
	c.AddFunc("30 14 * * *", CheckGuards)
	c.AddFunc("33 20 * * *", CheckGuards)
	c.Start()

	routes.Router()
}

func CheckGuards() {

	var guardShift string
	var advisoryGuard []models.AdvisorGuard
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
		"20", "33", "00")

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
		guardShift = "MATUTINO"
		repositories.New().GetAdvisoryGuardByShift(guardShift, morningGuardChange, &advisoryGuard)
	}

	if afternoonParse == afternoonLayoutParse {
		guardShift = "VESPERTINO"
		repositories.New().GetAdvisoryGuardByShift(guardShift, afternoonGuardChange, &advisoryGuard)
	}

	if eveningParse == eveningLayoutParse {
		guardShift = "NOCTURNO"
		response, err := repositories.New().GetAdvisoryGuardByShift(guardShift, eveningGuardChange, &advisoryGuard)
		if err != nil {
			return
		}
		for response.Next() {
			fmt.Println(response.Scan())
		}
	}

	fmt.Println("GUARDS:", advisoryGuard)
	for _, v := range advisoryGuard {
		advisorBitrix := models.AdvisorBitrix{
			UserID:        v.AdvisorBitrixID,
			PersonalSreet: "GUARDIA " + v.Development,
		}
		_ = models.UpdateBitrixGuardAdvisor(&advisorBitrix)
	}

	fmt.Println("GUARD SHIFT:", guardShift)
	fmt.Println("CURRENT DATE:", currentDate)
	fmt.Println("MORNING:", morningGuardChange, morningChangeLayout)
	fmt.Println("AFTERNOON:", afternoonGuardChange, afternoonChangeLayout)
	fmt.Println("EVENING:", eveningGuardChange, eveningChangeLayout)
}

package utils

import (
	"fmt"
	"time"
)

func ValidateCurrentDate(initDate, endDate, guardShift string) (bool, error) {

	layout := "2006-01-02 15:04:05"
	now := time.Now().UTC()

	currentTime, err := time.Parse(layout, fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()))
	if err != nil {
		return false, err
	}

	switch guardShift {
	case "MATUTINO":
		morningInit, err := time.Parse(layout, fmt.Sprintf("%v %02v:%02v:%02v",
			initDate, "10", "00", "00"))
		if err != nil {
			return false, err
		}

		morningEnd, err := time.Parse(layout, fmt.Sprintf("%v %02v:%02v:%02v",
			endDate, "14", "30", "00"))
		if err != nil {
			return false, err
		}

		if morningInit.Unix() >= currentTime.Unix() || morningEnd.Unix() <= currentTime.Unix() {
			return true, nil
		}
	case "VESPERTINO":
		afternoonInit, err := time.Parse(layout, fmt.Sprintf("%v %02v:%02v:%02v",
			initDate, "14", "30", "59"))
		if err != nil {
			return false, err
		}

		afternoonEnd, err := time.Parse(layout, fmt.Sprintf("%v %02v:%02v:%02v",
			endDate, "19", "00", "00"))
		if err != nil {
			return false, err
		}

		if afternoonInit.Unix() >= currentTime.Unix() && afternoonEnd.Unix() <= currentTime.Unix() {
			return true, nil
		}
	}

	return false, nil
}

func ValidateDates(initDate, endDate string) (err error) {

	layout := "2006-01-02"
	// Parse the date string into Go's time object
	// The 1st param specifies the format, 2nd is our date string
	init, err := time.Parse(layout, initDate)

	fmt.Println(init, err)

	if err != nil {
		return
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return
	}

	if init.Unix() > end.Unix() {
		err = fmt.Errorf("la fecha de inicio (%v) no puede ser mayor a la fecha final", init)
	}

	if end.Unix() < init.Unix() {
		err = fmt.Errorf("la fecha final (%v) no puede ser mayor a la fecha de inicio", end)
	}

	if init.Unix() == end.Unix() {
		return
	}
	return
}

func ValidateGuard(initDate, endDate string) (response bool, err error) {

	layout := "2006-01-02 15:04:05"
	now := time.Now().UTC()
	morningInitLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"10", "00", "00")
	morningEndLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"14", "30", "00")
	afternoonInitLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"14", "30", "59")
	afternoonEndLayout := fmt.Sprintf("%d-%02d-%02d %02v:%02v:%02v",
		now.Year(), now.Month(), now.Day(),
		"19", "00", "00")

	// Parse the date string into Go's time object
	// The 1st param specifies the format, 2nd is our date string
	init, err := time.Parse(layout, initDate)
	if err != nil {
		return
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return
	}

	morningInit, err := time.Parse(layout, morningInitLayout)
	if err != nil {
		return
	}
	morningEnd, err := time.Parse(layout, morningEndLayout)
	if err != nil {
		return
	}

	afternoonInit, err := time.Parse(layout, afternoonInitLayout)
	if err != nil {
		return
	}
	afternoonEnd, err := time.Parse(layout, afternoonEndLayout)
	if err != nil {
		return
	}

	if init.Unix() >= morningInit.Unix() && end.Unix() <= morningEnd.Unix() {
		return true, nil
	}

	if init.Unix() >= afternoonInit.Unix() && end.Unix() <= afternoonEnd.Unix() {
		return true, nil
	}
	return
}

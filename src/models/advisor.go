package models

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	ID           int
	BitrixID     string `json:"ID"`
	Name         string `json:"NAME"`
	LastName     string `json:"LAST_NAME"`
	Email        string `json:"EMAIL"`
	Photo        string `json:"PERSONAL_PHOTO"`
	WorkPosition string `json:"WORK_POSITION"`
	UserType     string `json:"USER_TYPE"`
	Active       bool   `json:"ACTIVE"`
}

type ResponseAdvisors struct {
	Result []BitrixAdvisors `json:"result"`
}

type BitrixAdvisors struct {
	ID           string `json:"ID"`
	Name         string `json:"NAME"`
	LastName     string `json:"LAST_NAME"`
	Email        string `json:"EMAIL"`
	Photo        string `json:"PERSONAL_PHOTO"`
	WorkPosition string `json:"WORK_POSITION"`
	UserType     string `json:"USER_TYPE"`
	Active       bool   `json:"ACTIVE"`
}

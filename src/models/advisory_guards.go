package models

import (
	"time"

	"gorm.io/gorm"
)

type AdvisorGuard struct {
	gorm.Model
	ID                  int
	AdvisorBitrixID     string
	Name                string
	LastName            string
	DevelopmentBitrixID string
	Development         string
	StartGuard          *time.Time
	EndGuard            *time.Time
}

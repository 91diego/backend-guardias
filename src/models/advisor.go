package models

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	ID       int
	BitrixID string
	Name     string
	Email    string
}

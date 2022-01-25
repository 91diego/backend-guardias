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

// CreateAdvisoryGuard
func CreateAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard) (err error) {
	err = db.Create(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdvisoryGuards
func GetAdvisoryGuards(db *gorm.DB, advisoryGuard *[]AdvisorGuard) (err error) {
	err = db.Find(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdvisoryGuardByID
func GetAdvisoryGuardByID(db *gorm.DB, advisoryGuard *AdvisorGuard, id string) (err error) {
	err = db.Where("id = ?", id).First(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateAdvisoryGuard
func UpdateAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard) (err error) {
	db.Save(advisoryGuard)
	return nil
}

// DeleteAdvisoryGuard
func DeleteAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard, id string) (err error) {
	db.Where("id = ?", id).Delete(advisoryGuard)
	return nil
}

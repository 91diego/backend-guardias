package models

import (
	"strings"
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

// GetAdvisoryGuards retrieve all advisory guards
func GetAdvisoryGuards(db *gorm.DB, advisoryGuard *[]AdvisorGuard) (err error) {
	err = db.Find(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdvisoryGuardByID retrieve by advisor id
func GetAdvisoryGuardByID(db *gorm.DB, advisoryGuard *AdvisorGuard, id string) (err error) {
	err = db.Where("id = ?", id).First(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateAdvisoryGuard create new advisory guard
func CreateAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard) (err error) {
	err = db.Create(advisoryGuard).Error
	if err != nil {
		return err
	}

	advisorBitrix := AdvisorBitrix{
		UserID:        advisoryGuard.AdvisorBitrixID,
		PersonalSreet: "GUARDIA " + advisoryGuard.Development,
	}
	// Update field on bitrix24 user profile
	UpdateBitrixGuardAdvisor(&advisorBitrix)
	return nil
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func UpdateAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard) (err error) {
	db.Save(advisoryGuard)
	advisorBitrix := AdvisorBitrix{
		UserID:        advisoryGuard.AdvisorBitrixID,
		PersonalSreet: "GUARDIA " + strings.ToUpper(advisoryGuard.Development),
	}
	// Update field on bitrix24 user profile
	UpdateBitrixGuardAdvisor(&advisorBitrix)
	return nil
}

// DeleteAdvisoryGuard delete advisory guard by id
func DeleteAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard, id string) (err error) {
	db.Where("id = ?", id).Delete(advisoryGuard)
	return nil
}

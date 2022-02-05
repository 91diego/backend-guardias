package models

import (
	"gorm.io/gorm"
)

type AdvisorGuard struct {
	gorm.Model
	ID                  int
	AdvisorBitrixID     string `json:"advisor_bitrix_id"`
	Name                string `json:"name"`
	LastName            string `json:"last_name"`
	DevelopmentBitrixID string `json:"development_bitrix_id"`
	Development         string `json:"development"`
	StartGuard          string `json:"start_guard"`
	EndGuard            string `json:"end_guard"`
	GuardShift          string `json:"guard_shift"`
}

// GetAdvisoryGuards retrieve all advisory guards
func GetAdvisoryGuards(db *gorm.DB, advisoryGuard *[]AdvisorGuard) (err error) {
	err = db.Find(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdvisoryGuardsParams retrieve advisory guards depending on params
func GetAdvisoryGuardsParams(db *gorm.DB, development, startDate, endDate string, advisoryGuard *[]AdvisorGuard) (err error) {

	var query string
	// all records on the date range
	if development == "" && startDate != "" && endDate != "" {
		query = "start_guard = ? and end_guard = ?"
		err = db.Where(query, startDate, endDate).Find(&advisoryGuard).Error
	}

	// records per development on the date range
	if development != "" && startDate != "" && endDate != "" {
		query = "development_bitrix_id = ? and start_guard BETWEEN ? AND ?"
		err = db.Where(query, development, startDate, endDate).Find(&advisoryGuard).Error
	}

	if err != nil {
		return err
	}
	return
}

// GetAdvisoryGuardByID retrieve by advisor id
func GetAdvisoryGuardByID(db *gorm.DB, advisoryGuard *AdvisorGuard, id string) (err error) {
	err = db.Where("id = ?", id).First(advisoryGuard).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdvisoryGuardByDate
func GetAdvisoryGuardByDate(db *gorm.DB, advisoryGuard *AdvisorGuard) (queryResult *gorm.DB, err error) {
	db.Begin()
	queryResult = db.Where("start_guard = ? and end_guard = ? and development_bitrix_id = ?",
		advisoryGuard.StartGuard, advisoryGuard.EndGuard, advisoryGuard.DevelopmentBitrixID).Find(&advisoryGuard)
	if queryResult.Error != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return queryResult, nil
}

// CreateAdvisoryGuard create new advisory guard
func CreateAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard) (err error) {
	db.Begin()
	err = db.Create(advisoryGuard).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func UpdateAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard) (err error) {
	db.Begin()
	err = db.Save(&advisoryGuard).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// DeleteAdvisoryGuard delete advisory guard by id
func DeleteAdvisoryGuard(db *gorm.DB, advisoryGuard *AdvisorGuard, id string) (err error) {
	db.Begin()
	err = db.Where("id = ?", id).Delete(advisoryGuard).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

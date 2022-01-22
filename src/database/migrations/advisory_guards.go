package migrations

import (
	"github.com/91diego/backend-guardias/src/database"
	"github.com/91diego/backend-guardias/src/models"
	"gorm.io/gorm"
)

type AdvisoryGuardRepo struct {
	Db *gorm.DB
}

func NewAdvisoryGuard() *AdvisoryGuardRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.AdvisorGuard{})
	return &AdvisoryGuardRepo{Db: db}
}

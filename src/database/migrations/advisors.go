package migrations

import (
	"github.com/91diego/backend-guardias/src/database"
	"github.com/91diego/backend-guardias/src/models"
	"gorm.io/gorm"
)

type AdvisorRepo struct {
	Db *gorm.DB
}

func NewAdvisor() *AdvisorRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Advisor{})
	return &AdvisorRepo{Db: db}
}

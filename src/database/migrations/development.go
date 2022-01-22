package migrations

import (
	"github.com/91diego/backend-guardias/src/database"
	"github.com/91diego/backend-guardias/src/models"
	"gorm.io/gorm"
)

type DevelopmentRepo struct {
	Db *gorm.DB
}

func NewDevelopment() *DevelopmentRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Development{})
	return &DevelopmentRepo{Db: db}
}

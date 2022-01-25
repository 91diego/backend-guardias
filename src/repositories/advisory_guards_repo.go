package repositories

import (
	"errors"
	"net/http"

	"github.com/91diego/backend-guardias/src/database"
	"github.com/91diego/backend-guardias/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdvisoryGuardRepo struct {
	Db *gorm.DB
}

func New() *AdvisoryGuardRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.AdvisorGuard{})
	return &AdvisoryGuardRepo{Db: db}
}

// GetAdvisoryGuards retrieve all advisory guards
func (repository *AdvisoryGuardRepo) GetAdvisoryGuards(c *gin.Context) {

	var advisoryGuard []models.AdvisorGuard

	err := models.GetAdvisoryGuards(repository.Db, &advisoryGuard)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de guardias.",
		"code":    http.StatusOK,
		"items":   advisoryGuard,
	})
}

// GetAdvisoryGuardByID retrieve by advisor id
func (repository *AdvisoryGuardRepo) GetAdvisoryGuardByID(c *gin.Context) {

	id, _ := c.Params.Get("id")
	var advisoryGuard models.AdvisorGuard

	err := models.GetAdvisoryGuardByID(repository.Db, &advisoryGuard, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Guardia listada exitosamente.",
		"code":    http.StatusOK,
		"items":   advisoryGuard,
	})
}

// CreateAdvisoryGuard create new advisory guard
func (repository *AdvisoryGuardRepo) CreateAdvisoryGuard(c *gin.Context) {

	var advisoryGuard models.AdvisorGuard

	c.BindJSON(&advisoryGuard)

	// TODO
	// Validate if advisory guard exists by date and development
	// If not, create de advisory guard
	// And update PERSONAL_STREET field on BITRIX24

	err := models.CreateAdvisoryGuard(repository.Db, &advisoryGuard)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "La guardia ha sido asignada exitosamente.",
		"code":    http.StatusCreated,
		"items":   advisoryGuard,
	})
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func (repository *AdvisoryGuardRepo) UpdateAdvisoryGuard(c *gin.Context) {

	var advisoryGuard models.AdvisorGuard
	id, _ := c.Params.Get("id")

	err := models.GetAdvisoryGuardByID(repository.Db, &advisoryGuard, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.BindJSON(&advisoryGuard)
	err = models.UpdateAdvisoryGuard(repository.Db, &advisoryGuard)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La guardia ha sido actualizada exitosamente.",
		"code":    http.StatusOK,
		"items":   advisoryGuard,
	})
}

// DeleteAdvisoryGuard delete advisory guard by id
func (repository *AdvisoryGuardRepo) DeleteAdvisoryGuard(c *gin.Context) {

	var advisoryGuard models.AdvisorGuard
	id, _ := c.Params.Get("id")

	err := models.DeleteAdvisoryGuard(repository.Db, &advisoryGuard, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusGone, gin.H{
		"message": "La guardia ha sido eliminada.",
		"code":    http.StatusGone,
		"items":   advisoryGuard,
	})
}

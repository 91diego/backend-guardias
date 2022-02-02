package repositories

import (
	"errors"
	"fmt"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de guardias.",
		"code":    http.StatusOK,
		"status":  "success",
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
			c.JSON(http.StatusNotFound, gin.H{
				"message": "La guardia no existe.",
				"code":    http.StatusNotFound,
				"status":  "warning",
				"items":   "",
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Guardia listada exitosamente.",
		"code":    http.StatusOK,
		"status":  "success",
		"items":   advisoryGuard,
	})
}

// CreateAdvisoryGuard create new advisory guard
func (repository *AdvisoryGuardRepo) CreateAdvisoryGuard(c *gin.Context) {

	// TODO
	// VALIDATE DATES, START DATE CAN NOT BE GREATHER THAN END DATE
	// AND END DATE CAN NOT BE LOWER THAN START DATE
	var rows int
	var advisoryGuard models.AdvisorGuard
	c.BindJSON(&advisoryGuard)
	res, err := models.GetAdvisoryGuardByDate(repository.Db, &advisoryGuard)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	rows = int(res.RowsAffected)
	if rows < 1 {
		err = models.CreateAdvisoryGuard(repository.Db, &advisoryGuard)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
				"code":    http.StatusInternalServerError,
				"status":  "warning",
				"items":   "",
			})
		}
		// Update field on bitrix24 user profile
		advisorBitrix := models.AdvisorBitrix{
			UserID:        advisoryGuard.AdvisorBitrixID,
			PersonalSreet: "GUARDIA " + advisoryGuard.Development,
		}
		models.UpdateBitrixGuardAdvisor(&advisorBitrix)
		c.JSON(http.StatusCreated, gin.H{
			"message": "La guardia ha sido asignada exitosamente.",
			"code":    http.StatusCreated,
			"status":  "success",
			"items":   advisoryGuard,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"message": fmt.Sprintf("La fecha %v se encuentra ocupada por el asesor %v %v. ¿Desea actualizar la guardia?",
				advisoryGuard.StartGuard, advisoryGuard.Name, advisoryGuard.LastName),
			"code":   http.StatusFound,
			"status": "warning",
			"items":  advisoryGuard,
		})
	}
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func (repository *AdvisoryGuardRepo) UpdateAdvisoryGuard(c *gin.Context) {

	// TODO
	// VALIDATE DATES, START DATE CAN NOT BE GREATHER THAN END DATE
	// AND END DATE CAN NOT BE LOWER THAN START DATE

	var advisoryGuard models.AdvisorGuard
	id := c.Query("id")

	err := models.GetAdvisoryGuardByID(repository.Db, &advisoryGuard, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "La guardia no existe.",
				"code":    http.StatusNotFound,
				"status":  "warning",
				"items":   "",
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
	}

	c.BindJSON(&advisoryGuard)
	err = models.UpdateAdvisoryGuard(repository.Db, &advisoryGuard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
	}

	// Update field on bitrix24 user profile
	advisorBitrix := models.AdvisorBitrix{
		UserID:        advisoryGuard.AdvisorBitrixID,
		PersonalSreet: "GUARDIA " + advisoryGuard.Development,
	}
	models.UpdateBitrixGuardAdvisor(&advisorBitrix)

	c.JSON(http.StatusOK, gin.H{
		"message": "La guardia ha sido modificada.",
		"code":    http.StatusOK,
		"status":  "success",
		"items":   advisoryGuard,
	})
}

// DeleteAdvisoryGuard delete advisory guard by id
func (repository *AdvisoryGuardRepo) DeleteAdvisoryGuard(c *gin.Context) {

	var advisoryGuard models.AdvisorGuard
	id := c.Query("id")

	err := models.GetAdvisoryGuardByID(repository.Db, &advisoryGuard, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "La guardia no existe.",
				"code":    http.StatusNotFound,
				"status":  "warning",
				"items":   "",
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
	}

	err = models.DeleteAdvisoryGuard(repository.Db, &advisoryGuard, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
	}
	c.JSON(http.StatusGone, gin.H{
		"message": "La guardia ha sido eliminada.",
		"code":    http.StatusGone,
		"status":  "success",
		"items":   "",
	})
}

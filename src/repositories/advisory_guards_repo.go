package repositories

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/91diego/backend-guardias/src/database"
	"github.com/91diego/backend-guardias/src/models"
	"github.com/91diego/backend-guardias/src/utils"
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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de guardias.",
		"code":    http.StatusOK,
		"status":  "success",
		"items":   advisoryGuard,
	})
}

// GetAdvisoryGuardsParams retrieve advisory guards depending on params
func (repository *AdvisoryGuardRepo) GetAdvisoryGuardsParams(c *gin.Context) {

	var advisoryGuard []models.AdvisorGuard
	development := c.Query("development")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "La fecha de inicio y la fecha final son necesarios.",
			"code":    http.StatusBadRequest,
			"status":  "warning",
			"items":   "",
		})
		return
	}

	err := models.GetAdvisoryGuardsParams(repository.Db, development, startDate, endDate, &advisoryGuard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
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
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Guardia listada exitosamente.",
		"code":    http.StatusOK,
		"status":  "success",
		"items":   advisoryGuard,
	})
}

func (repository *AdvisoryGuardRepo) CheckAdvisoryGuardByDate(c *gin.Context) {

	var rows int
	var advisoryGuard models.AdvisorGuard
	c.BindJSON(&advisoryGuard)

	err := utils.ValidateDates(advisoryGuard.StartGuard, advisoryGuard.EndGuard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
	}

	res, err := models.CheckAdvisoryGuardByDate(repository.Db, &advisoryGuard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error: %v.", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
	}
	rows = int(res.RowsAffected)
	if rows > 0 {
		c.JSON(http.StatusFound, gin.H{
			"message": fmt.Sprintf("La fecha %v se encuentra ocupada por el asesor %v %v. ¿Desea actualizar la guardia?",
				advisoryGuard.StartGuard, advisoryGuard.Name, advisoryGuard.LastName),
			"code":   http.StatusFound,
			"status": "warning",
			"items":  advisoryGuard,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"code":    http.StatusOK,
		"status":  "success",
		"items":   "",
	})
}

// CreateAdvisoryGuard create new advisory guard
func (repository *AdvisoryGuardRepo) CreateAdvisoryGuard(c *gin.Context) {

	var advisoryGuard models.AdvisorGuard
	c.BindJSON(&advisoryGuard)

	// Update field on bitrix24 user profile
	advisorBitrix := models.AdvisorBitrix{
		UserID:        advisoryGuard.AdvisorBitrixID,
		PersonalSreet: "GUARDIA " + advisoryGuard.Development,
	}

	currentDate, err := utils.ValidateCurrentDate(advisoryGuard.StartGuard, advisoryGuard.EndGuard, advisoryGuard.GuardShift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"items":   "",
		})
		return
	}

	// Add guard field to bitrix24 user
	if currentDate {
		err := models.UpdateBitrixGuardAdvisor(&advisorBitrix)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
				"code":    http.StatusInternalServerError,
				"status":  "error",
				"items":   "",
			})
			return
		}
	}

	// Create record on DB
	err = models.CreateAdvisoryGuard(repository.Db, &advisoryGuard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error: %v.", err),
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"items":   "",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "La guardia ha sido asignada exitosamente.",
		"code":    http.StatusCreated,
		"status":  "success",
		"items":   advisoryGuard,
	})

}

// UpdateAdvisoryGuard update advisory guard by advisor id
func (repository *AdvisoryGuardRepo) UpdateAdvisoryGuard(c *gin.Context) {

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
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
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
		return
	}

	currentDate, err := utils.ValidateCurrentDate(advisoryGuard.StartGuard, advisoryGuard.EndGuard, advisoryGuard.GuardShift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"items":   "",
		})
		return
	}
	// Update field on bitrix24 user profile
	if currentDate {
		advisorBitrix := models.AdvisorBitrix{
			UserID:        advisoryGuard.AdvisorBitrixID,
			PersonalSreet: "GUARDIA " + advisoryGuard.Development,
		}
		err := models.UpdateBitrixGuardAdvisor(&advisorBitrix)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
				"code":    http.StatusInternalServerError,
				"status":  "error",
				"items":   "",
			})
			return
		}
	}

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
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
	}

	currentDate, err := utils.ValidateCurrentDate(advisoryGuard.StartGuard, advisoryGuard.EndGuard, advisoryGuard.GuardShift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
			"code":    http.StatusInternalServerError,
			"status":  "error",
			"items":   "",
		})
		return
	}

	err = models.DeleteAdvisoryGuard(repository.Db, &advisoryGuard, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Ha ocurrido un error %v: ", err),
			"code":    http.StatusInternalServerError,
			"status":  "warning",
			"items":   "",
		})
		return
	}

	// Update field on bitrix24 user profile
	if currentDate {
		advisorBitrix := models.AdvisorBitrix{
			UserID:        advisoryGuard.AdvisorBitrixID,
			PersonalSreet: "",
		}
		err := models.UpdateBitrixGuardAdvisor(&advisorBitrix)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Ha ocurrido un error: %v ", err),
				"code":    http.StatusInternalServerError,
				"status":  "error",
				"items":   "",
			})
			return
		}
	}
	c.JSON(http.StatusGone, gin.H{
		"message": "La guardia ha sido eliminada.",
		"code":    http.StatusGone,
		"status":  "success",
		"items":   "",
	})
}

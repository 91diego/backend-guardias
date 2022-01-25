package advisoryguards

import (
	"github.com/91diego/backend-guardias/src/repositories"
	"github.com/gin-gonic/gin"
)

// GetAllAdvisoryGuards retrieve all advisory guards
func GetAdvisoryGuards(c *gin.Context) {
	repositories.New().GetAdvisoryGuards(c)
}

// GetAdvisoryGuardByID retrieve by advisor id
func GetAdvisoryGuardByID(c *gin.Context) {
	repositories.New().GetAdvisoryGuardByID(c)
}

// CreateAdvisoryGuard Create new advisory guard
func CreateAdvisoryGuard(c *gin.Context) {
	repositories.New().CreateAdvisoryGuard(c)
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func UpdateAdvisoryGuard(c *gin.Context) {
	repositories.New().UpdateAdvisoryGuard(c)
}

// DeleteAdvisoryGuard delete advisory guard by id
func DeleteAdvisoryGuard(c *gin.Context) {
	repositories.New().DeleteAdvisoryGuard(c)
}

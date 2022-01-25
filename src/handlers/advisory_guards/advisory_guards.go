package advisoryguards

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAdvisoryGuards retrieve all advisory guards
func GetAdvisoryGuards(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetAdvisoryGuards METHOD",
		"code":    http.StatusOK,
		"items":   "",
	})
}

// GetAdvisoryGuardByID retrieve by advisor id
func GetAdvisoryGuardByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetAdvisoryGuardByID METHOD",
		"code":    http.StatusOK,
		"items":   "",
	})
}

// CreateAdvisoryGuard Create new advisory guard
func CreateAdvisoryGuard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CreateAdvisoryGuard METHOD",
		"code":    http.StatusOK,
		"items":   "",
	})
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func UpdateAdvisoryGuard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateAdvisoryGuard METHOD",
		"code":    http.StatusOK,
		"items":   "",
	})
}

// DeleteAdvisoryGuard delete advisory guard by id
func DeleteAdvisoryGuard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteAdvisoryGuard METHOD",
		"code":    http.StatusOK,
		"items":   "",
	})
}

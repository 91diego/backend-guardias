package advisoryguards

import "github.com/gin-gonic/gin"

// GetAllAdvisoryGuards retrieve all advisory guards
func GetAdvisoryGuards(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAdvisoryGuards METHOD",
		"code":    200,
		"items":   "",
	})
}

// GetAdvisoryGuard retrieve by advisor id
func GetAdvisoryGuardByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAdvisoryGuardByID METHOD",
		"code":    200,
		"items":   "",
	})
}

// CreateAdvisoryGuard Create new advisory guard
func CreateAdvisoryGuard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateAdvisoryGuard METHOD",
		"code":    200,
		"items":   "",
	})
}

// UpdateAdvisoryGuard update advisory guard by advisor id
func UpdateAdvisoryGuard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateAdvisoryGuard METHOD",
		"code":    200,
		"items":   "",
	})
}

// DeleteAdvisoryGuard delete advisory guard by id
func DeleteAdvisoryGuard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteAdvisoryGuard METHOD",
		"code":    200,
		"items":   "",
	})
}

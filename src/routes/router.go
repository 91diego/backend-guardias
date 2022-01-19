package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/91diego/backend-guardias/src/handlers/bitrix24/advisors"
)

func Router() {
	router := gin.Default()

	advisorsV1 := router.Group("api/v1/advisors")
	{
		advisorsV1.GET("", advisors.GetAdvisors)
	}

	/*userV1 := router.Group("api/v1/user")
	{
		userV1.GET("", users.GetUsers)
		userV1.GET("/by-id", users.GetUserById)
		userV1.POST("", users.CreateUser)
		userV1.PUT("", users.UpdateUser)
		userV1.DELETE("", users.DeleteUser)
	}*/

	router.Run()
}

package routes

import (
	"github.com/gin-gonic/gin"

	migration_advisor "github.com/91diego/backend-guardias/src/database/migrations"
	"github.com/91diego/backend-guardias/src/handlers/bitrix24/advisors"
	"github.com/91diego/backend-guardias/src/handlers/bitrix24/developments"
)

func Router() {
	router := gin.Default()

	advisorsV1 := router.Group("api/v1/advisors")
	{
		migration_advisor.NewAdvisor()
		advisorsV1.GET("", advisors.GetAdvisors)
	}

	developmentV1 := router.Group("api/v1/developments")
	{
		developmentV1.GET("", developments.GetDevelopments)
	}

	router.Run()
}

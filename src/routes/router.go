package routes

import (
	"github.com/gin-gonic/gin"

	migrations "github.com/91diego/backend-guardias/src/database/migrations"
	advisoryguards "github.com/91diego/backend-guardias/src/handlers/advisory_guards"
	"github.com/91diego/backend-guardias/src/handlers/bitrix24/advisors"
	"github.com/91diego/backend-guardias/src/handlers/bitrix24/developments"
)

func Router() {
	router := gin.Default()

	advisorsV1 := router.Group("api/v1/advisors")
	{
		migrations.NewAdvisor()
		advisorsV1.GET("", advisors.GetAdvisors)
	}

	developmentV1 := router.Group("api/v1/developments")
	{
		developmentV1.GET("", developments.GetDevelopments)
	}

	advisoryGuardsV1 := router.Group("api/v1/advisory-guard")
	{
		migrations.NewAdvisoryGuard()
		advisoryGuardsV1.GET("", advisoryguards.GetAdvisoryGuards)
		advisoryGuardsV1.GET("/params", advisoryguards.GetAdvisoryGuardsParams)
		advisoryGuardsV1.GET("/by-id", advisoryguards.GetAdvisoryGuardByID)
		advisoryGuardsV1.POST("", advisoryguards.CreateAdvisoryGuard)
		advisoryGuardsV1.PUT("", advisoryguards.UpdateAdvisoryGuard)
		advisoryGuardsV1.DELETE("", advisoryguards.DeleteAdvisoryGuard)
	}

	router.Run()
}

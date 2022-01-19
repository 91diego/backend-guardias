package main

import (
	"github.com/91diego/backend-guardias/src/routes"
	"github.com/91diego/backend-guardias/src/utils"
)

func main() {

	utils.EnvVariables()
	routes.Router()
}

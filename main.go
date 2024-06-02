package main

import (
	initializers "github.com/NabinGrz/SocialMedia/src/config"
	databaseService "github.com/NabinGrz/SocialMedia/src/models/database"
	"github.com/NabinGrz/SocialMedia/src/router"
)

func init() {
	initializers.LoadEnvVariables()
	databaseService.DBConnection()
}
func main() {
	r := router.Router(databaseService.DB)

	r.Run()
}

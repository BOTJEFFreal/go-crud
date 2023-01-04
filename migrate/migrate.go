package main

import (
	"example/main/databasePackage"
	"example/main/models"
)

func init() {
	databasePackage.LoadEnvVariables()
	databasePackage.ConnectToDB()
}
func main() {
	databasePackage.DB.AutoMigrate(&models.Post{})
}

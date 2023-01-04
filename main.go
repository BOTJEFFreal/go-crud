package main

import (
	"example/main/controllers"
	"example/main/databasePackage"

	"github.com/gin-gonic/gin"
)

func init() {
	databasePackage.LoadEnvVariables()
	databasePackage.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts/")
	r.GET("/posts/:startDate/:endDate/", controllers.PostsShow)
	r.GET("/posts/", controllers.PostsAll)
	r.Run()
}

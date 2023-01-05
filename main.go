package main

import (
	"example/main/controllers"
	"example/main/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts/", controllers.PostsAdd)
	r.GET("/posts/:startDate/:endDate/", controllers.PostsShow)
	r.GET("/posts/", controllers.PostsAll)
	r.Run()
}

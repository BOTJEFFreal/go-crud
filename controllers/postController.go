package controllers

import (
	"example/main/initializers"
	"example/main/models"

	"github.com/gin-gonic/gin"
)

func PostsAdd(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
}

func PostsAll(c *gin.Context) {

	//Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id
	start := c.Param("startDate")
	end := c.Param("endDate")

	//Get the post
	var posts []models.Post
	initializers.DB.Where("created_at BETWEEN ? AND ?", start, end).Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

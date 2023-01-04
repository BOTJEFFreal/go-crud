package controllers

import (
	"example/main/databasePackage"
	"example/main/models"
	"time"

	"github.com/gin-gonic/gin"
)

func PostsAdd(c *gin.Context) {
	var body struct {
		Title string
		Body  string
		Date  time.Time
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body, Date: body.Date}
	result := databasePackage.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
}

func PostsAll(c *gin.Context) {

	//Get the post
	var posts []models.Post
	databasePackage.DB.Find(&posts)

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
	databasePackage.DB.Where("title BETWEEN ? AND ?", start, end).Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

package controllers

import (
	"github.com/adamsyah24/test-crud/initializers"
	"github.com/adamsyah24/test-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Ambil data req body
	var body struct {
		Title   string
		Content string
		Body    string
	}

	c.Bind(&body)

	//Buat post

	post := models.Post{Title: body.Title, Content: body.Content, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Ambil data post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Ambil ID dari url
	id := c.Param("id")

	//Ambil data post
	var post models.Post
	initializers.DB.First(&post, id)

	//Response
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//Ambil ID dari url
	id := c.Param("id")

	//Ambil data dari req body
	var body struct {
		Title   string
		Content string
		Body    string
	}

	c.Bind(&body)

	//Cari post yang diupdate
	var post models.Post
	initializers.DB.First(&post, id)

	//Update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title:   body.Title,
		Content: body.Content,
		Body:    body.Body})

	//REsponse
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Ambil ID dari url
	id := c.Param("id")

	//Cari post yang diupdate
	initializers.DB.Delete(&models.Post{}, id)

	//response
	c.Status(200)
}

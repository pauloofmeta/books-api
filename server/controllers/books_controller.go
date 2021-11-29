package controllers

import (
	"books-api/database"
	"books-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context)  {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Books
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context)  {
	db := database.GetDatabase()
	var book models.Books
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(201, book)
}

func ListBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Books
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context)  {
	db := database.GetDatabase()
	var book models.Books
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context)  {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Books{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(200)
}
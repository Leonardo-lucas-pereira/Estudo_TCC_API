package controllers

import (
	"strconv"

	"github.com/Leonardo-lucas-pereira/tcc-api/database"
	"github.com/Leonardo-lucas-pereira/tcc-api/models"
	"github.com/Leonardo-lucas-pereira/tcc-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var user models.User

	err = db.First(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user" + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func ListUsers(c *gin.Context) {
	db := database.GetDatabase()
	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list users: " + err.Error(),
		})
		return
	}
	c.JSON(200, users)
}

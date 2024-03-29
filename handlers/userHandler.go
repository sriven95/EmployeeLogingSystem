package handlers

import (
	"net/http"

	"Employee-Log-System/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AllUser(c *gin.Context) {
	var re []models.User
	h.DB.Find(&re)
	c.IndentedJSON(http.StatusOK, re)
}

func (h handler) AddUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}

func (h handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	err := h.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted Successfully",
	})
}

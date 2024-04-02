package users

import (
	"fmt"
	"go-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("created user")
	sqlClient := c.MustGet("sqlClient").(*gorm.DB)
	manager := models.NewUserManager(sqlClient)
	if !manager.Create(user) {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "user could not be created"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

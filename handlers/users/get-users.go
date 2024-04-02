package users

import (
	"go-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	sqlClient := c.MustGet("sqlClient").(*gorm.DB)

	manager := models.NewUserManager(sqlClient)
	res, err := manager.GetMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "user could not be created"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

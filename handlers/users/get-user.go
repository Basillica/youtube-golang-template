package users

import (
	"go-template/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	sqlClient := c.MustGet("sqlClient").(*gorm.DB)
	u64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	manager := models.NewUserManager(sqlClient)
	res, err := manager.GetById(uint(u64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "user could not be created"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

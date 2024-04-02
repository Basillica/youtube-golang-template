package users

import (
	"go-template/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	sqlClient := c.MustGet("sqlClient").(*gorm.DB)
	manager := models.NewUserManager(sqlClient)

	u64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	if err := manager.Delete(uint(u64)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "user could not be deleted:" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
}

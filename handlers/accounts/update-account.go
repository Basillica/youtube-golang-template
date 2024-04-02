package accounts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	fmt.Println("updated account")
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

package accounts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	fmt.Println("deleted account")
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

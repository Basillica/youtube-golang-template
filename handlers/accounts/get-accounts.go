package accounts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context) {
	fmt.Println("gotten accounts")
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

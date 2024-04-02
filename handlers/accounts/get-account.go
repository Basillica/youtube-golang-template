package accounts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Account
// @Summary      Endpoint to get account info
// @Description  Get account info using id, user ID and token
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        id	path	string  true  "id"
// @Param        user_oid	path	string  true  "User ID"
// @Param        access_token	path	string  true  "Access Token"
// @Success      200  {object}  responses.Account
// @Router       /account/:id [get]
func GetAccount(c *gin.Context) {
	fmt.Println("gotten account")
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

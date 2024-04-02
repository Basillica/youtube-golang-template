package accounts

import (
	"fmt"
	"go-template/types/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Account
// @Summary      Endpoint to create a account
// @Description  create a account with the token and user id
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        req	body	responses.Account	true	"Account Request"
// @Param        user_oid	path	string  true  "User ID"
// @Param        access_token	path	string  true  "Access Token"
// @Success      201  {object}  responses.Account
// @Failure      500  {object}  errors.ErrorMessage
// @Router       /create [post]
func CreateAccount(c *gin.Context) {
	if true {
		c.JSON(http.StatusUnauthorized, &errors.ErrorMessage{
			Error:      "an error occured retrieving user from the database. The user might not exist yet",
			StatusCode: http.StatusUnauthorized,
			Message:    "UNAUTHORIZED",
		})
		return
	}
	
	fmt.Println("created account")
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "optional":
		return ""
	case "email":
		return "This field must be a valid email address"
	case "oneof":
		return "This field must be one of 'SUPER ADMIN' 'JO ADMIN' 'JO SERVICE' 'CUSTOMER ADMIN' 'CUSTOMER USER' 'CUSTOMER OPERATOR'"
	case "required":
		return "This field is required field and has to be provided"
	case "alpha":
		return "This field should contain only english alphabets"
	case "lte":
		return "This field should be less than " + fe.Param()
	case "gte":
		return "This field should be greater than " + fe.Param()
	case "min":
		return "This field should not be less than " + fe.Param()
	case "max":
		return "This field should not be greater than " + fe.Param()
	}
	return "Unknown error"
}

func FormatError(c *gin.Context, err error) {
	fmt.Println(err, "the frigging error")
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{
				Field:   fe.Field(),
				Message: GetErrorMsg(fe)}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

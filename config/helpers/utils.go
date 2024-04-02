package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnOnErr(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
}

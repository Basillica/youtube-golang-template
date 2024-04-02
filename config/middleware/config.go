package middleware

import (
	"fmt"
	"net/http"

	"go-template/config/helpers"

	"github.com/gin-gonic/gin"
)

func ConfigMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appConfig, err := helpers.LoadConfig(".")
		if err != nil {
			fmt.Println(err)
			c.Abort()
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.Set("appConfig", &appConfig)
	}
}

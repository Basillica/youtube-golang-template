package misc

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	fmt.Println(c.Writer.Header())
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

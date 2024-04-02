package middleware

import (
	cfg "go-template/types/config"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	alllowedHeaders := "Content-Type, Content-Length, Accept-Encoding, User-Scope, User-Role, Authorization, accept, origin, Cache-Control, X-Requested-With, content-disposition"
	return func(c *gin.Context) {
		ae := c.MustGet("appConfig").(*cfg.AppConfig)
		c.Writer.Header().Set("Access-Control-Allow-Origin", ae.CORS_ORIGIN)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", alllowedHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

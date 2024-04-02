package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "go-template/docs"
)

func AddSwaggerRoutes(r *gin.Engine, v string) {
	g := r.Group(v + "/swagger")
	g.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

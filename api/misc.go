package api

import (
	"go-template/handlers/misc"

	"github.com/gin-gonic/gin"
)

func AddMiscRoutes(r *gin.Engine, v string) {
	a := r.Group(v + "/misc")
	a.GET("/health", misc.Health)
}

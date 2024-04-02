package api

import (
	"go-template/config/middleware"
	acc "go-template/handlers/accounts"

	"github.com/gin-gonic/gin"
)

func AddAccountsRoutes(r *gin.Engine, v string) {
	a := r.Group(v + "/accounts")
	a.Use(middleware.AuthMiddleware())
	a.GET("/accs", acc.GetAccounts)
	a.GET("/accs/:id", acc.GetAccount)
	a.POST("/create", acc.CreateAccount)
	a.PUT("/accs/:id", acc.UpdateAccount)
	a.DELETE("/accs/delete/:id", acc.DeleteAccount)
}

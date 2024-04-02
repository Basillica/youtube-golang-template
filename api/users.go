package api

import (
	"go-template/config/middleware"
	"go-template/handlers/users"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine, v string) {
	a := r.Group(v + "/users")
	a.Use(middleware.AuthMiddleware(), middleware.SQLMiddleware())
	a.GET("/users", users.GetUsers)
	a.GET("/user/:id", users.GetUser)
	a.POST("/create", users.CreateUser)
	a.PUT("/user/:id", users.UpdateUser)
	a.DELETE("/user/delete/:id", users.DeleteUser)
}

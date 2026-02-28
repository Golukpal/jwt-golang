package routes

import (
	"github.com/Golukpal/jwt/cmd/controllers"
	"github.com/Golukpal/jwt/cmd/routes"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine){
	user:= routes.Group("/users")
	user.POST("/signup", controllers.Signup())
	user.POST("/login", controllers.Login())
}
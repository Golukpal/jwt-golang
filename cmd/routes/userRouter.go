package routes

import (
	"github.com/Golukpal/jwt/cmd/middleware"
	"github.com/Golukpal/jwt/cmd/routes"
	"github.com/gin-gonic/gin"
	"github.com/Golukpal/jwt/cmd/controllers"
)

func UserRoutes(routes *gin.Engine){
	routes.Use(middleware.Authenticate())
	routes.GET("/users", controllers.GetUsers())
	routes.GET("/users/:user_id", controllers.GetUser())

}
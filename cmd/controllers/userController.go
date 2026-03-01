package controllers

import (
	"net/http"

	"github.com/Golukpal/cmd/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"github.com/Golukpal/cmd/helpers"
)

var usercollection *mongo.Collection = db.OpenCollection(db.Client, "user")
var validates = validator.New()

func HashPassword()

func varifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		 userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c,userId); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}
 

	}
}

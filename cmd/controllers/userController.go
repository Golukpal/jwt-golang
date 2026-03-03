package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Golukpal/cmd/db"
	"github.com/Golukpal/cmd/helpers"
	"github.com/Golukpal/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
 

	}
}

package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Golukpal/cmd/db"
	"github.com/Golukpal/cmd/helpers"
	"github.com/Golukpal/cmd/models"
	"github.com/Golukpal/jwt/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")
var validates = validator.New()

func HashPassword()

func varifyPassword()

func Signup()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err:= c.BindJSON(&user); err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}

		validationErr:= validate.Struct(user)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err!= nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for email"})
			return 
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err!= nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for phone"})
			return 
		}

		}
}

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

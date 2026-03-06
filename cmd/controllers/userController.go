package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Golukpal/cmd/db"
	"github.com/Golukpal/cmd/helpers"
	"github.com/Golukpal/cmd/models"
	"github.com/Golukpal/jwt/cmd/helpers"
	"github.com/Golukpal/jwt/cmd/models"
	_ "github.com/Golukpal/jwt/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")
var validate = validator.New()

func HashPassword()

func VarifyPassword(userPassword string, provudedPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true 
	msg := ""

	if err!= nil{
		msg = fmt.Sprintf("email and password is incorrect")
		check = false 
	}
	return check, msg
	  
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for email"})
			return
		}

		// count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		// defer cancel()
		// if err!= nil{
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for phone"})
		// 	return
		// }

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email and phone already exist"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = premitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helpers.GenrateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("user item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
		}
		passwordIsValid, msg := VarifyPassword(*user.Passwaord, *&foundUser.Password)
		defer cancel()
		if passwordIsValid!= true{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return 
		}

		if foundUser.Email == nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})

		}
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Emial, *&foundUser.First_name, *foundUser.Last_name, *foundUser.User_id, *foundUser.User_type)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err := userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
			return
		}
		c.JSON(http.StatusOK, foundUser)




	}
}

func GetUsers()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)

	}
}

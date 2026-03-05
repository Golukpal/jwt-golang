package helpers

import (
	"context"
	"error"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Golukpal/cmd/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)
type SignedDetails struct {
	Email    string
	First_name string
	Last_name string
	Uid      string
	User_type string
	jwt.standardClaims 
}

var useCollection *mongo.Collection = db.OpenCollection(db.Client, "user")
 
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, uid string, userType string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:    email,
		First_name: firstName,
		Last_name: lastName,
		Uid:      uid,
		User_type: userType,
		standardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
},
	}
	refreshClaims := &SignedDetails{
		standardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}	
	return token, refreshToken, err
		}

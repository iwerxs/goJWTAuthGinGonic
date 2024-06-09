package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iwerxs/goJWTAuthGinGonic/database"
	helper "github.com/iwerxs/goJWTAuthGinGonic/helpers"
	"github.com/iwerxs/goJWTAuthGinGonic/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// "fmt"
// "log"
// "time"
// "os"
// "context"
// "strconv"
// "github.com/gin-gonic/gin"
// "github.com/go-playground/validator/v10"
// helper "goJWTAuthGinGonic/helpers"
// "goJWTAuthGinGonic/models"
// "goJWTAuthGinGonic/helpers"
// "golang.org/x/crypto/bcrypt"

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var validate = validator.New()

// setup functions defined in userRouter
func HashPassword()

func VerifyPassword()

func Register()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"500 error": err.Error()})
			return
		}
	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userID := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"400 error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		// Decode json data type into readable golang data
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"500 error": err.Error()})
			return
		}
		c.JSON(http.StatusOk, user)
	}
}
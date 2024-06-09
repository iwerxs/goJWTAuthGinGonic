package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iwerxs/goJWTAuthGinGonic/database"
	helper "github.com/iwerxs/goJWTAuthGinGonic/helpers"
	"github.com/iwerxs/goJWTAuthGinGonic/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
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
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"500 error":validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"500 error":"error occcured while checking email"})
		}
		// count, err := userCollection.CountDocuments(ctx, bson.M{"company":user.Company})
		// defer cancel()
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"500 error":"error occcured while checking company"})
		// }

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email already exists"})
		}

		// create User Object
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		// insert User into database
		resultInsertNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("user profile no created")
			c.JSON(http.StatusInternalServerError, gin.H{"500 error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertNumber)
	}
}

// Login function
func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
	}
}

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
		err := userCollection.FindOne(ctx, bson.M{"user_id":userID}).Decode(&user)
		// Decode json data type into readable golang data
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"500 error": err.Error()})
			return
		}
		c.JSON(http.StatusOk, user)
	}
}
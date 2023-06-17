package controllers

import (
	"authServer/database"
	"authServer/helpers"
	"authServer/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
func HashPassword(){

}

func VerifyPassword(){

}

func Signup() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		
		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}

		validationError := validate.Struct(user)
		if validationError != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}
		count,err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Error checking email"})
		}
		count,err = userCollection.CountDocuments(ctx, bson.M{"phone_number":user.Phone})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Error checking Phone No"})
		}
		// SO we a checking if the phon no ar email already exists in database 
		// this is working by checking is any one of both count assigned changes value then this user might already exist
		if count > 0 {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"This email or phone no already exist"})
		}
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		user.Token = token
		user.Refresh_token = &refreshToken
		resultInsertNumber, insertError := userCollection.InsertOne(ctx, user)
		insertError != nil{
			msg := fmt.Sprintf("User was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,resultInsertNumber)
	}

}
func Login(){

}
func GetUsers(){

}
func GetUserByID() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// setting time out 
		var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// creating a variable to store the incoming query user
		var user models.User
		// finding the user in collection
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,user)
	}
}

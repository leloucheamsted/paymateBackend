package main

import (
	"context"
	"log"
	services "paymate/Services"
	"paymate/middleware"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const valName = "USER_UID"

var key = ""

const (
	credFileName = "./credentials/serviceAccountKey.json"
)

var app *firebase.App
var firestoreClient *firestore.Client

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		c.Header("Exposed", "Content-Length, Content-Range")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	ctx := context.Background()
	log.Println(services.TestPayment())
	services.TestPaymentStatus()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	middleware, err := middleware.NewMiddleware(app)
	if err != nil {
		panic(err)
	}
	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	r.Use(CORSMiddleware())
	r.Use(middleware.MiddlewareFunc())
	//r.GET("/user", users.CreateUser(app))
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

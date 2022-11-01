package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"paymate/handler"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
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
	//cards.TestGetCardHolderDetails()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer firestoreClient.Close()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Authorization", "token"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "token"},
		AllowCredentials: true,

		AllowAllOrigins: true,
		MaxAge:          8640024576987,
	}))

	r.Use(CORSMiddleware())
	// middleware, err := newMiddleware()
	// if err != nil {
	// 	panic(err)
	// }
	//r.GET("/verify/:token", middleware.verifyToken)
	//r.Use(middleware.MiddlewareFunc())
	r.POST("/User/Create", handler.CreateUser)
	r.GET("/User/GetUser/:id", handler.GetUserByID)
	r.GET("/GetUser/AllTransactions/:id", handler.GetAllTransaction)
	r.GET("/GetUser/LatestTransactions/:id", handler.GetLatestTransaction)
	r.GET("/GetUser/ListCards/:id", handler.GetUserCards)
	r.PUT("/User/Update/:id", handler.UpdateUser)
	r.POST("/Wallet/Reload", handler.ReloadWallet)
	r.GET("/Wallet/Reload/Confirm/:id", handler.ConfirmReload)
	r.POST("/Card/Create", handler.CreateCard)
	r.POST("/Card/Reload", handler.ReloadCard)
	r.GET("/Card/Details/:card_id", handler.CardDetails)
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func (fam *FirebaseAuthMiddleware) verifyToken(c *gin.Context) {

	token := c.Param("token")
	fmt.Println(token)

	_, err := fam.cli.VerifyIDToken(context.Background(), token)

	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Token is valid")
	c.JSON(http.StatusOK, gin.H{"error": "", "message": "token valid"})

}

type FirebaseAuthMiddleware struct {
	cli *auth.Client
}

func newMiddleware() (*FirebaseAuthMiddleware, error) {

	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return &FirebaseAuthMiddleware{
		cli: auth,
	}, nil
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": code, "message": message})
}
func (fam *FirebaseAuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := fam.cli.VerifyIDToken(context.Background(), token)

		if err != nil {
			RespondWithError(c, http.StatusUnauthorized, "Invalid API token")
			return
		}

		log.Println("User ID is " + idToken.UID)
		//c.Set(valName, idToken.UID)
		c.Next()
	}
}

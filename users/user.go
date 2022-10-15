package users

import (
	"context"
	"log"
	"net/http"
	"paymate/users/services"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const (
	credFileName = "./credentials/serviceAccountKey.json"
)

var app *firebase.App
var firestoreClient *firestore.Client

type User struct {
	uiud      string `firebase:"id"`
	firstName string `firebase:"firstName"`
	LastName  string `firebase:"lastName"`
	amount    string `firebase:"amount"`
	phone     string `firebase:"phone"`
}

func CreateUser(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	services.NewUser(app, c)
	c.JSON(http.StatusOK, gin.H{"user": "user"})
}

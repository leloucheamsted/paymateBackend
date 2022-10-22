package handler

import (
	"context"
	"log"
	"paymate/user"

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

func CreateUser(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	newUser := new(user.Users)
	newUser = &user.Users{
		Uiud:      c.PostForm("uiud"),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Amount:    0,
		Phone:     c.PostForm("phone"),
	}
	user.AddUser(app, c, *newUser)
}
func GetUserByID(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	id := c.Param("id")
	log.Println(id)
	user.GetUser(app, c, id)
}
func UpdateUser(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	User := new(user.Users)
	User = &user.Users{
		Uiud:      c.Param("id"),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
	}
	log.Println(User)
	user.UpdateUser(app, c, *User)
}

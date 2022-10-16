package handler

import (
	"context"
	"log"
	"paymate/user"
	"strconv"

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
	_amount, err := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	newUser = &user.Users{
		Uiud:      c.PostForm("uiud"),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Amount:    _amount,
		Phone:     c.PostForm("phone"),
	}
	user.AddUser(app, c, *newUser)
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
		Uiud:      c.Param("user_id"),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
	}
	log.Println(User)
	user.UpdateUser(app, c, *User)
}

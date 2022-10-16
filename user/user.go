package user

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
)

type Users struct {
	Uiud      string `firebase:"id,omitempty"`
	FirstName string `firebase:"firstName,omitempty"`
	LastName  string `firebase:"lastName,omitempty"`
	Amount    int64  `firebase:"amount,omitempty"`
	Phone     string `firebase:"phone,omitempty"`
}

func AddUser(app *firebase.App, c *gin.Context, user Users) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)
	defer client.Close()

	log.Println(user)

	var use = map[string]interface{}{
		"uiud":      user.Uiud,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"amount":    user.Amount,
		"phone":     user.Phone,
	}

	newUser, err := client.Collection("Users").Doc(string(user.Uiud)).Set(ctx, use)
	if err != nil {
		panic(err)
	}
	log.Println(newUser)
	c.JSON(200, gin.H{"message": "User added with success", "user": use})
}

func UpdateUser(app *firebase.App, c *gin.Context, user Users) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)
	defer client.Close()

	log.Println(user)

	var use = map[string]interface{}{
		"uiud":      user.Uiud,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	}

	User := client.Collection("Users").Doc((string(user.Uiud)))
	_, err = User.Update(ctx, []firestore.Update{
		{Path: "firstName", Value: use["firstName"]},
		{Path: "lastName", Value: use["lastName"]},
	})
	userData := client.Collection("Users").Doc(string(user.Uiud))
	doc, err := userData.Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(doc.Data())
	c.JSON(200, gin.H{"message": "user profil was be  updated with success", "user": doc.Data()})
}

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
	if err != nil {
		log.Println(err)
	}
	userData := client.Collection("Users").Doc(string(user.Uiud))
	doc, err := userData.Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(doc.Data())
	c.JSON(200, gin.H{"message": "user profil was be  updated with success", "user": doc.Data()})
}

func ReloadAmountUser(app *firebase.App, userId string, amount float64) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	updateAmount := client.Collection("Users").Doc(userId) // try to acess document
	_, err = updateAmount.Update(ctx, []firestore.Update{
		{Path: "amount", Value: firestore.Increment(amount)},
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("UpdateAmount=>  Amount Update succesfully")
	}
}

func RemoveAmountUser(app *firebase.App, userId string, amount float64) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	updateAmount := client.Collection("Users").Doc(userId) // try to acess document
	_, err = updateAmount.Update(ctx, []firestore.Update{
		{Path: "amount", Value: firestore.Increment(-amount)},
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("UpdateAmount=>  Amount Update succesfully")
	}
}

func UpdateUserCardHolder(app *firebase.App, userId string, params map[string]interface{}) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	updateAmount := client.Collection("Users").Doc(userId) // try to acess document
	_, err = updateAmount.Update(ctx, []firestore.Update{
		{Path: "cardHolder", Value: params},
	})
	if err != nil {
		log.Println("Error=> upload card holder details failed")
		log.Println(err)
	} else {
		log.Println("UpdateCardHolder=>  Card details Update succesfully")
	}
}

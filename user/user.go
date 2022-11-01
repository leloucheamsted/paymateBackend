package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"

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
		"createAt":  time.Now().Format("2006-01-02 15:04:05"),
	}

	newUser, err := client.Collection("Users").Doc(string(user.Uiud)).Set(ctx, use)
	if err != nil {
		panic(err)
	}
	log.Println(newUser)
	c.JSON(200, gin.H{"message": "User added with success", "user": use})
}

func GetUser(app *firebase.App, c *gin.Context, id string) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	defer client.Close()
	var user Users
	log.Println(id)

	userData, err := client.Collection("Users").Doc(id).Get(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println(userData.Data())
	//user.FirstName = userData.Data()["firstName"].(string)
	//user.LastName = userData.Data()["lastName"].(string)
	//user.Amount = int64(userData.Data()["amount"].(int64))
	//user.Phone = userData.Data()["phone"].(string)
	fmt.Printf("Document data: %#v\n", user)
	if err != nil {
		log.Println(err)
		c.JSON(202, gin.H{"status": "failed", "message": err, "data": userData.Data()})
	}
	log.Println(user)
	c.JSON(200, gin.H{"status": "success", "user": userData.Data()})

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
	log.Println("OnData=> ", userId, amount)
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

func RemoveAmountUser(app *firebase.App, userId string, amount float64) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()
	log.Println("OnData => RemoveAmountUser ", userId, amount)
	updateAmount := client.Collection("Users").Doc(userId) // try to acess document
	_, err = updateAmount.Update(ctx, []firestore.Update{
		{Path: "amount", Value: firestore.Increment(-amount)},
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("OnResponse =>  Amount Update succesfully")
	}
}

func UpdateUserCardHolder(app *firebase.App, userId string, params map[string]interface{}) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	params["createAt"] = time.Now().Format("2006-01-02 15:04:05")
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

func GetLatestTransaction(app *firebase.App, c *gin.Context, Id string) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	transactions := []map[string]interface{}{}
	getTransactionLimit := client.Collection("Transactions").OrderBy("createAt", firestore.Desc).Where("userId", "==", Id).Limit(15).Documents(ctx)
	fmt.Println(getTransactionLimit.GetAll())
	if getTransactionLimit != nil {

		for {
			doc, err := getTransactionLimit.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				//log.Fatalf("Failed to iterate: %v", err)
				c.JSON(200, transactions)
				return
			} else {
				transaction := doc.Data()
				transactions = append(transactions, transaction)

			}
		}
		c.JSON(200, transactions)

	} else {
		c.JSON(200, transactions)

	}

}
func GetCardsUser(app *firebase.App, c *gin.Context, Id string) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	var cards []map[string]interface{}
	getAllCards := client.Collection("Cards").Where("data.meta_data.user_id", "==", Id).Documents(ctx)
	if getAllCards != nil {
		for {
			doc, err := getAllCards.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalln("Failed to iterate: #{err}")
			}
			card := doc.Data()
			cards = append(cards, card)
		}
		c.JSON(200, gin.H{"message": "List of  cards return succesfully", "status": "success", "cards": cards})
	} else {
		c.JSON(200, gin.H{"message": "User don't have transactions", "status": "success", "cards": cards})
	}
}

func GetAllTransaction(app *firebase.App, c *gin.Context, Id string) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	var transactions []map[string]interface{}
	getTransactionLimit := client.Collection("Transactions").OrderBy("createAt", firestore.Desc).Where("userId", "==", Id).Documents(ctx)
	if getTransactionLimit != nil {

		for {
			doc, err := getTransactionLimit.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			transaction := doc.Data()
			transactions = append(transactions, transaction)
		}
		c.JSON(200, gin.H{"message": "List of  transactions return succesfully", "status": "success", "transactions": transactions})
	} else {
		c.JSON(200, gin.H{"message": "User don't have transactions", "status": "success", "transactions": transactions})

	}

}

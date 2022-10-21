package transaction

import (
	"context"
	"log"
	"paymate/user"

	firebase "firebase.google.com/go"
)

type Users struct {
	Uiud      string `firebase:"id,omitempty"`
	FirstName string `firebase:"firstName,omitempty"`
	LastName  string `firebase:"lastName,omitempty"`
	Amount    int64  `firebase:"amount,omitempty"`
	Phone     string `firebase:"phone,omitempty"`
}

func AddTransaction(app *firebase.App, transaction map[string]interface{}) {
	log.Println(transaction)
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	newTransaction, err := client.Collection("Transactions").NewDoc().Create(ctx, transaction)
	if err != nil {
		log.Println(err)
	}
	log.Println(transaction["id"])
	user.ReloadAmountUser(app, transaction["userId"].(string), transaction["amount"].(float64))
	log.Println("New transaction=>", newTransaction)
}

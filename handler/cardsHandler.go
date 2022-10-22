package handler

import (
	"context"
	"log"
	"paymate/cards"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func CreateCard(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	newHolder := new(cards.CardHolder)
	newHolder = &cards.CardHolder{
		UserID:    c.PostForm("userId"),
		FirstName: c.PostForm("firstName"),
		LastName:  c.PostForm("lastName"),
		Phone:     c.PostForm("phone"),
	}
	log.Println("CardHolder=>", newHolder)
	var response1 = cards.RegisterCardHolder(*newHolder)
	var data map[string]interface{}

	if response1["status"] == "success" {
		c.JSON(200, gin.H{"status": "success", "data": response1})
	} else {
		c.JSON(401, gin.H{"status": "failed", "data": data})
	}

}

func ReloadCard(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	newFund := new(cards.FunCard)
	newFund = &cards.FunCard{
		CardId: c.PostForm("card_id"),
		Amount: c.PostForm("amount"),
		UserId: c.PostForm("user_Id"),
	}
	log.Println("FundCard=>", newFund)
	var response1 = cards.ReloadCard(*newFund)
	var data map[string]interface{}
	if response1["status"] == "success" {
		var response = cards.GetCardDetails(newFund.CardId)
		c.JSON(200, gin.H{"status": "success", "data": response})
	} else {
		c.JSON(401, gin.H{"status": "failed", "data": data})
	}
}

func CardDetails(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	CardId := c.Param("card_id")
	var response1 = cards.GetCardDetails(CardId)
	var data map[string]interface{}
	if response1["status"] == "success" {
		c.JSON(200, gin.H{"status": "success", "data": response1})
	} else {
		c.JSON(401, gin.H{"status": "failed", "data": data})
	}
}

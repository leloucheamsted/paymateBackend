package users

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

type User struct {
	uiud      string `firebase:"id"`
	firstName string `firebase:"firstName"`
	fastName  string `firebase:"lastName"`
	amount    string `firebase:"amount"`
	phone     string `firebase:"phone"`
}

func CreateUser(app *firebase.App, c *gin.Context) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"user": "user"})
}

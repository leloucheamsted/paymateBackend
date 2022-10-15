package services

import (
	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
)

type User struct {
	uiud      string `firebase:"id"`
	firstName string `firebase:"firstName"`
	lastName  string `firebase:"lastName"`
	amount    int64  `firebase:"amount"`
	phone     string `firebase:"phone"`
}

func NewUser(app *firebase.App, c *gin.Context) {
	// ctx := context.Background()
	// client, err := app.Firestore(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// u := uuid.NewV4()
	// user := &User{
	// 	uiud:      u.String(),
	// 	firstName: "",
	// 	lastName:  "",
	// 	amount:    100,
	// 	phone:     "",
	// }
	// fmt.Println(user)
	// defer client.Close()
	// newUser, err := client.Collection("Users").NewDoc().Create(ctx, user)

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(newUser)
	c.JSON(200, gin.H{"message": "User added with success", "data": "data"})
}

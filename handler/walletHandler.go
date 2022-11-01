package handler

import (
	"context"
	"log"
	"math/rand"
	"paymate/user"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var (
	url             = "https://sandbox.fapshi.com"
	apikey          = "FAK_TEST_c056736b7a6e7ef836b2"
	apiuser         = "a264a3b6-58b7-45ff-ba51-35cda4f24623"
	alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
)

type Payment struct {
	Amount     int
	ExternalId string
	UserId     string
	Phone      string
	Message    string
	Email      string
}
type PaymentResponse struct {
	Message string
	TransId string
	Date    string
}

func RandomString(n int, alphabet []rune) string {

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
func ReloadWallet(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	ExtId := RandomString(20, alphabet)
	_amount, err := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	if err != nil {
		log.Println(err)
	}

	newPayment := Payment{
		Amount:     int(_amount),
		ExternalId: ExtId,
		UserId:     c.PostForm("userId"),
		Phone:      c.PostForm("phone"),
		Message:    "Hello Paymate",
		Email:      "cabrauleketchanga@gmail.com",
	}
	var payment = PaymentFunc(newPayment)

	c.JSON(201, gin.H{"message": "Payment Accepted", "data": payment})

}

func ConfirmReload(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	transId := c.Param("id")
	c.JSON(200, PaymentStatusFunc(transId))
}

func GetLatestTransaction(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	Id := c.Param("id")

	user.GetLatestTransaction(app, c, Id)
}

func GetAllTransaction(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	Id := c.Param("id")

	user.GetAllTransaction(app, c, Id)
}

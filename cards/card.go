package cards

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"paymate/user"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	guuid "github.com/google/uuid"
	"google.golang.org/api/option"
)

var (
	number   []rune = []rune("123456789009876543210123456789")
	alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

	defaultUrl = "https://issuecards.api.bridgecard.co/v1/issuing/sandbox"
	token      = "at_test_23a31fec7d9296a8e26429662043fb59d3aadda239614b2c88107fe8dc4d0a0faf916f29aa25e1fe36d36703cac95fce266420f5c24db2403ba4c0c3662153c5419a1d637d5d1c4059b2b194483a1917f27ca39cc80689976cc5b2bc0f5b295ba7f688ce19ad161345637cb83d95133dd8eba5bd8dfd0fde0ad196ac5d466f20bc34600b961f06ce8529f0fea9c41886381c888f6aee1e95f283ac9433024ec296f64d42c1737bfa6855d83553ec5d3e00a55b8f0acf2fc809c63e9a931ab1bb924fbda52b0f5accb8bcf904652260da76970ae5a5240319ac2371c16c2fe83c5d1974ef2c8e956b934034d388efec9885521551b707027df3fa0458c43a2d47"
)

type CardHolder struct {
	FirstName string
	LastName  string
	UserID    string
	Phone     string
}

type FunCard struct {
	CardId string
	Amount string
	UserId string
}

const (
	device       = 657
	credFileName = "./credentials/serviceAccountKey.json"
)

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
func RandomInteger(n int, number []rune) string {

	alphabetSize := len(number)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := number[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

var app *firebase.App

func RegisterCardHolder(cardHolder CardHolder) map[string]interface{} {
	url := defaultUrl + "/cardholder/register_cardholder_synchronously"
	ctx := context.Background()
	opt := option.WithCredentialsFile(credFileName)
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	id := guuid.New()

	fmt.Println(id)
	method := "POST"
	bodyRequest := "{\n  \"first_name\":\" " + cardHolder.FirstName + "\",\n  \"last_name\":\" " + cardHolder.LastName + "\",\n  \"address\": {\n    \"address\": \"9 South Street\",\n    \"city\": \"Douala\",\n    \"state\": \"Littoral\",\n    \"country\": \"Cameroon\",\n    \"postal_code\": \"1000242\",\n    \"house_no\": \"13\"\n  },\n  \"phone\":\" " + cardHolder.Phone + "\",\n  \"email_address\":\"omnipay@gmail.com\",\n  \"identity\": {\n      \"id_type\": \"CAMEROON_NATIONAL_ID\",\n      \"id_no\":\" " + id.String() + "\",\n\"id_image\": \"https://image.com\",\n \"first_name\":\" " + cardHolder.FirstName + "\",\n\"last_name\":\"" + cardHolder.LastName + "\"\n  },\n   \"meta_data\":{\"userId\":\"" + cardHolder.UserID + "\"}\n}"
	payload := strings.NewReader(bodyRequest)

	var data map[string]interface{}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("token", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(body))
	json.Unmarshal(body, &data)
	var response map[string]interface{}
	holderId := data["data"].(map[string]interface{})["cardholder_id"]

	if data["status"] == "success" {
		responseStatut := GetCardHolderDetails(app, holderId.(string), cardHolder.UserID)
		if responseStatut["status"] == "success" {
			log.Println("OnResponse=> CardHolder Created")
			response = CreateUserCard(holderId.(string), cardHolder.UserID)
		}
	}
	return response
}

func GetCardHolderDetails(app *firebase.App, cardHolderId, user_id string) map[string]interface{} {
	cardHolder_id := cardHolderId
	url := defaultUrl + "/cardholder/get_cardholder?cardholder_id=" + cardHolder_id
	method := "GET"
	var data map[string]interface{}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("token", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(body))
	json.Unmarshal(body, &data)
	user.UpdateUserCardHolder(app, user_id, data)
	return data
}

func GetCardDetails(card_id string) map[string]interface{} {
	url := "https://issuecards-api-bridgecard-co.relay.evervault.com/v1/issuing/sandbox/cards/get_card_details?card_id=" + card_id
	var data map[string]interface{}

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("token", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	log.Println("OnResponse=> Get Card details")
	//data["createAt"] = time.Now().Format("2006-01-02 15:04:05")
	json.Unmarshal(body, &data)
	SaveCard(data)
	return data

}
func CreateUserCard(holderId, userId string) map[string]interface{} {
	url := defaultUrl + "/cards/create_card"
	var data map[string]interface{}
	var response map[string]interface{}
	method := "POST"

	bodyRequest := "{\n   \"cardholder_id\":\"" + holderId + "\",\n  \"card_type\": \"virtual\",\n  \"card_brand\": \"Visa\",\n  \"card_currency\": \"USD\",\n  \"meta_data\": {\"user_id\":\" " + userId + "\"}\n}"
	payload := strings.NewReader(bodyRequest)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("token", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	json.Unmarshal(body, &data)

	if data["status"] == "success" {
		log.Println("OnResponse=> Card Created")
		response = GetCardDetails(data["data"].(map[string]interface{})["card_id"].(string))
	}
	return response
}

func SaveCard(cards map[string]interface{}) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cards)
	defer client.Close()

	log.Println("OnResponse=> Saved card data operation")
	log.Println(cards)
	newCard, err := client.Collection("Cards").NewDoc().Create(ctx, cards)
	if err != nil {
		log.Println(err)
	}
	log.Println(newCard)
}

func ReloadCard(cardFunc FunCard) map[string]interface{} {
	url := defaultUrl + "/cards/fund_card"
	method := "PATCH"
	TransactionRef := RandomString(10, alphabet)
	_amount, err := strconv.Atoi(cardFunc.Amount)
	if err != nil {
		log.Println("OnError=> ", err)
	}
	_amountConv := _amount / device

	bodyRequest := "{\n    \"card_id\": " + cardFunc.CardId + ",\n  \"amount\": " + string(rune(_amountConv)) + ",\n  \"transaction_reference\": " + TransactionRef + ",\n  \"currency\": \"USD\"\n}"
	payload := strings.NewReader(bodyRequest)
	var data map[string]interface{}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("token", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Parse error=>", err)
	}
	if data["status"] == "success" {

		user.RemoveAmountUser(app, cardFunc.UserId, float64(_amount))
	}
	fmt.Println(data)
	return data
}

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	url             = "https://sandbox.fapshi.com"
	apikey          = "FAK_TEST_c056736b7a6e7ef836b2"
	apiuser         = "a264a3b6-58b7-45ff-ba51-35cda4f24623"
	alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
)

type Payment struct {
	amount     int
	externalId string
	userId     string
	phone      string
	message    string
}

func randomString(n int, alphabet []rune) string {

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

// Test direct payment
// Parameters we need
// Phone, amount,
func TestPayment() string {
	rand.Seed(time.Now().UnixNano())

	method := "POST"
	rs := randomString(20, alphabet)
	var data map[string]interface{}
	var ms = map[string]interface{}{
		"amount":     100,
		"externalId": rs,
		"phone":      "682421795",
		"userId":     "123",
		"message":    " Hello paymate",
	}
	jsonValue, err := json.Marshal(ms)
	fmt.Println(bytes.NewBuffer(jsonValue))
	client := &http.Client{}
	req, err := http.NewRequest(method, url+"/direct-pay", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apiuser", apiuser)
	req.Header.Add("apikey", apikey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &data)

	strValue := fmt.Sprint(data["transId"])

	return strValue
}

// Test payment status
// Parameter we need
// transaction Id
func TestPaymentStatus() {
	method := "GET"
	var data map[string]interface{}

	client := &http.Client{}
	req, err := http.NewRequest(method, url+"/payment-status/"+"2escfbquA0", nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("apiuser", apiuser)
	req.Header.Add("apikey", apikey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {

	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &data)

	fmt.Println(data)

}

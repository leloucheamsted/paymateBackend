package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	transaction "paymate/transactions"
	"strconv"
	"strings"
)

// Test direct payment
// Parameters we need
// Phone, amount,
func TestPayment(payment Payment) map[string]interface{} {
	method := "POST"
	var data map[string]interface{}

	userData := "{\n		\"amount\":      " + strconv.Itoa(payment.Amount) + " ,\n		\"email\":      \"" + payment.Email + "\",\n		\"externalId\": \"" + payment.ExternalId + "\",\n		\"userId\":     \"" + payment.UserId + "\",\n		\"message\":    \"" + payment.Message + "\",\n		\"phone\":      \"" + payment.Phone + "\"\n\n}"

	fmt.Println(userData)

	client := &http.Client{}
	req, err := http.NewRequest(method, url+"/direct-pay", strings.NewReader(userData))

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
		log.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &data)

	strValue := fmt.Sprint(data)
	log.Println("DirectPaymateResponse=>" + strValue)
	if res.StatusCode == 200 {
		data["statut"] = 200
	}
	return data
}

// Test payment status
// Parameter we need
// transaction Id
func TestPaymentStatus(transId string) map[string]interface{} {
	method := "GET"
	var data map[string]interface{}

	client := &http.Client{}
	req, err := http.NewRequest(method, url+"/payment-status/"+transId, nil)

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
		log.Println(err)
	}
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(body, &data)
		data["statut"] = "200"
		fmt.Println(data)
		if data["status"] == "SUCCESSFUL" {
			transaction.AddTransaction(app, data)
		}
		return data
	} else {
		return data
	}

}

package cards

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	defaultUrl = "https://issuecards.api.bridgecard.co/v1/issuing/sandbox"
	token      = "at_test_23a31fec7d9296a8e26429662043fb59d3aadda239614b2c88107fe8dc4d0a0faf916f29aa25e1fe36d36703cac95fce266420f5c24db2403ba4c0c3662153c5419a1d637d5d1c4059b2b194483a1917f27ca39cc80689976cc5b2bc0f5b295ba7f688ce19ad161345637cb83d95133dd8eba5bd8dfd0fde0ad196ac5d466f20bc34600b961f06ce8529f0fea9c41886381c888f6aee1e95f283ac9433024ec296f64d42c1737bfa6855d83553ec5d3e00a55b8f0acf2fc809c63e9a931ab1bb924fbda52b0f5accb8bcf904652260da76970ae5a5240319ac2371c16c2fe83c5d1974ef2c8e956b934034d388efec9885521551b707027df3fa0458c43a2d47"
)

func TestRegisterCardHolder() {
	url := defaultUrl + "/cardholder/register_cardholder_synchronously"
	method := "POST"

	payload := strings.NewReader(`{
		"first_name": "John",
		"last_name": "Doe",
		"address": {
		  "address": "9 South Street",
		  "city": "South",
		  "state": "South",
		  "country": "Cameroon",
		  "postal_code": "1000242",
		  "house_no": "13"
		},
		"phone": "676563421",
		"email_address":"testingboy@gmail.com",
		"identity": {
			"id_type": "CAMEROON_NATIONAL_ID",
			"id_no": "74838920202",
			"id_image": "https://image.com",
			"first_name": "John",
			"last_name": "Doe"
		},
		 "meta_data":{"any_key": "any_value"}
	  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("token", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestGetCardHolderDetails() {
	cardHolder_id := "4391b2fac2534b549a79b9411e6098ee"
	url := defaultUrl + "/cardholder/get_cardholder?cardholder_id=" + cardHolder_id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("token", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestCreateCard() {
	url := defaultUrl + "/cards/create_card"
	method := "POST"
	//cardHolder_id := "4391b2fac2534b549a79b9411e6098ee"
	payload := strings.NewReader(`{
	 "cardholder_id": "4391b2fac2534b549a79b9411e6098ee",
	"card_type": "virtual",
	"card_brand": "Visa",
	"card_currency": "USD",
	"meta_data": {"user_id": "d0658fedf828420786e4a7083fa"}
  }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("token", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
func TestGetDetailsCard() {

}

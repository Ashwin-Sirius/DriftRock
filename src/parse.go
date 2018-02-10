package main

/**
Go version:1.9.3

@Feb 2018
@by Ashwin Rana
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//local struct to store users from driftrock API
type UserData struct {
	Data []struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
	} `json:"data"`
}

//local struct to store purchase history from the API
type PurchaseData struct {
	Data []struct {
		UserID string `json:"user_id"`
		Item   string `json:"item"`
		Spend  string `json:"spend"`
	} `json:"data"`
}

//json user data parser
func getUsers() {
	url := fmt.Sprintf("https://driftrock-dev-test.herokuapp.com/users?page=1&per_page=50")

	client := &http.Client{}

	//build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var recordUser UserData

	if err := json.NewDecoder(resp.Body).Decode(&recordUser); err != nil {
		log.Println(err) //decode vs unmarshall (use decode when io is streaming from a source ie http and marshal when json is locally available)
	}

	fmt.Println(resp.Status)
	fmt.Println((recordUser))
}

//json purchase data parser
func getPurchases() {
	url := fmt.Sprintf("https://driftrock-dev-test.herokuapp.com/purchases?page=1&per_page=50")

	client := &http.Client{}

	//build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var recordPurchases PurchaseData

	if err := json.NewDecoder(resp.Body).Decode(&recordPurchases); err != nil {
		log.Println(err) //decode vs unmarshall (use decode when io is streaming from a source ie http and marshal when json is locally available)
	}

	fmt.Println(resp.Status)
	fmt.Println((recordPurchases))
}

// func enumerator() {
//// implement loop until Content-Length = 11 (end of content)
// }

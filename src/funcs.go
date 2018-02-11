package main

import "fmt"

/**
Go version:1.9.3

@Feb 2018
@by Ashwin Rana
*/

func most_Sold() {
	store := getPurchases().Data
	// var mostSoldItem string (algorithm to find the majority element)
	// maxCount := 0
	// maxIndex := 0
	// count := 1
	for i := 1; i < len(store); i++ {
		// fmt.Println(store[i].Item)
		// if store[i-1] == store[i] {
		// 	count++
		// }
		// if count > maxCount {
		// 	maxIndex = i
		// 	maxCount = count
		// }
		fmt.Println(store[i].Item)
	}
}

func total_Spend(email string) {
	// store:= getPurchases().Data
	// store2:= getUsers().Data

}

func most_Loyal() {
	// store:= getPurchases().Data
	// store2:= getUsers().Data

}

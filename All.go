package main

import (
	"fmt"
	"github.com/BelphegorPrime/lib"
)

func All(crypto float64) []string {
	returnArray := []string{}
	returnArray = append(returnArray, "Url         UserName         Password         Timestamp")
	rows, err := db.Query("SELECT * FROM website")
	if err != nil {
		fmt.Println("can't execute select query: " + err.Error())
	}
	for rows.Next() {
		w := new(Website)
		err := rows.Scan(&w.ID, &w.Url, &w.UserName, &w.Password, &w.Timestamp)
		if err != nil {
			fmt.Println("can't read into struct: " + err.Error())
		}
		returnArray = append(returnArray, w.ToString())
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Error with Row: " + err.Error())
	}

	if crypto >= 0 {
		for i := 0; i < len(returnArray); i++ {
			cipherText, err := lib.Encrypt([]byte(returnArray[i]), key)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
			returnArray[i] = string(cipherText)
		}
	}

	return returnArray
}

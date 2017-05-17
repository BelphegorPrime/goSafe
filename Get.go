package main

import (
	"fmt"
)

func Get(unDecryptedUrl string) []string {
	url, err := decrypt([]byte(unDecryptedUrl))
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	returnArray := []string{}
	if len(url) != 0 {
		returnArray = append(returnArray, "Url         UserName         Password         Timestamp")
		rows, err := db.Query("SELECT * FROM website WHERE url like ?", url)
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
	} else {
		returnArray = append(returnArray, "no get parameter given")
	}
	return returnArray
}

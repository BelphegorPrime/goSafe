package main

import (
	"fmt"
)

func Get(requestContent map[string]interface{}) []byte {
	unDecryptedUrl := requestContent["url"].(string)
	url, err := decrypt([]byte(unDecryptedUrl))
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	if len(url) != 0 {
		var returnString string = "Url         UserName         Password         Timestamp"
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
			returnString = returnString + w.ToString()
		}
		err = rows.Err()
		if err != nil {
			fmt.Println("Error with Row: " + err.Error())
		}

		return []byte(returnString)
	} else {
		return []byte("no get parameter given")
	}
}

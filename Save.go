package main

import (
	"fmt"
	"time"
)

func Save(url string, username string, password string) []byte {

	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
		count, errorstring := getCount(url, username, password)
		if(errorstring != nil){
			return errorstring
		}
		if(count == 0){
			_, err := db.Exec("INSERT INTO website("+
				"url, "+
				"username, "+
				"password, "+
				"timestamp) "+
				"VALUES(?, ?, ?, ?);",
				url,
				username,
				password,
				time.Now(),
			)
			if err != nil {
				fmt.Println("Can not insert data into Database: " + err.Error())
				return []byte("Can not insert data into Database")
			} else {
				return []byte("Everything worked fine!")
			}
		} else {
			return []byte("Entry already exists.")
		}


	} else {
		return []byte("not enough parameters given")
	}
}

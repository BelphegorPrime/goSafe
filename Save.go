package main

import (
	"fmt"
	"time"
)

func Save(requestContent map[string]interface{}) []byte {
	url := requestContent["url"].(string)
	username := requestContent["username"].(string)
	password := requestContent["password"].(string)
	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
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
		return []byte("not enough parameters given")
	}
}

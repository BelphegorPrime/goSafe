package main

import (
	"fmt"
	"net/http"
	"time"
)

func save_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
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
			fmt.Println("Can't insert data into Database: " + err.Error())
			ciphertext, err := encrypt([]byte("Can not insert data into Database"))
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
			rw.Write(ciphertext)
		} else {
			ciphertext, err := encrypt([]byte("Everything worked fine!"))
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}
			rw.Write(ciphertext)
		}

	} else {
		ciphertext, err := encrypt([]byte("not enough parameters given"))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rw.Write(ciphertext)
	}
}

func get_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)

	undecryptedUrl := requestContent["url"].(string)
	url, err := decrypt([]byte(undecryptedUrl))
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

		ciphertext, err := encrypt([]byte(returnString))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rw.Write(ciphertext)
	} else {
		ciphertext, err := encrypt([]byte("no get parameter given"))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		rw.Write(ciphertext)
	}
}

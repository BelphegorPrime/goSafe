package main

import (
	"fmt"
)

func getCount(url string, username string, password string) (int, []byte) {
	var count int
	rows, err := db.Query("SELECT count(*) "+
		"FROM website "+
		"WHERE url like ? AND "+
		"username like ? AND "+
		"password like ?",
		url, username, password)
	if err != nil {
		fmt.Println("Can not Count the amount of Elements: " + err.Error())
		return 0, []byte("Can not Count the amount of Elements.")
	}
	if rows.Next() {
		rows.Scan(&count)
	}
	return count, nil
}

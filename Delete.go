package main

import (
	"fmt"
)

func Delete(url string, username string, password string) []byte {
	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
		var count uint
		rows, err := db.Query("SELECT count(*) " +
			"FROM website " +
			"WHERE url like ? AND " +
			"username like ? AND " +
			"password like ?",
			url, username, password)
		if err != nil {
			fmt.Println("Can not find anything to delete: " + err.Error())
			return []byte("Theres nothing to delete.")
		}
		if rows.Next() {
			rows.Scan(&count)
		}
		if(count > 0){
			_, err = db.Exec("DELETE FROM website "+
				"WHERE url like ? AND " +
				"username like ? AND " +
				"password like ?",
				url, username, password)
			if err != nil {
				fmt.Println("Can not delete anything from database: " + err.Error())
				return []byte("Can not delete anything from database.")
			}
			return []byte("Everything worked fine!")
		}else{
			return []byte("Theres nothing to delete.")
		}
	} else {
		return []byte("not enough parameters given")
	}
}

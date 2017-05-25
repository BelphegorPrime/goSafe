package main

import (
	"fmt"
	"strconv"
)

func Delete(url string, username string, password string) []byte {
	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
		count, errorstring := getCount(url, username, password)
		if errorstring != nil {
			return errorstring
		}
		if count > 0 {
			_, err := db.Exec("DELETE FROM website "+
				"WHERE url like ? AND "+
				"username like ? AND "+
				"password like ?",
				url, username, password)
			if err != nil {
				fmt.Println("Can not delete anything from database: " + err.Error())
				return []byte("Can not delete anything from database.")
			}
			return []byte("Everything worked fine! " + strconv.Itoa(count) + " entries were deleted")
		} else {
			return []byte("Theres nothing to delete.")
		}
	} else {
		return []byte("not enough parameters given")
	}
}

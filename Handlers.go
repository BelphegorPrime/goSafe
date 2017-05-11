package main
import (
	"net/http"
	"fmt"
	"runtime"
	"bytes"
	"encoding/json"
)

type Website struct {
	ID		int
	Url     	string
	UserName     	string
	Password 	string
}

func (w Website) ToString() string {
	return w.Url+" "+w.UserName+" "+w.Password
}

func getRequestContentFromRequest(req *http.Request) map[string]interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	data := buf.Bytes()
	var requestContent map[string]interface{}
	err := json.Unmarshal(data, &requestContent)
	if err != nil {
		fmt.Println(err)
	}
	return requestContent
}

func index_func(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}
func save_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	fmt.Println(requestContent)

	url := requestContent["url"].(string)
	username := requestContent["username"].(string)
	password := requestContent["password"].(string)
	if len(url) != 0 && len(username) != 0 && len(password) != 0 {
		_, err := db.Exec("INSERT INTO website("+
			"url, "+
			"username, "+
			"password) "+
			"VALUES(?, ?, ?);",
			url,
			username,
			password,
			)
		if err != nil {
			fmt.Println("Can't insert data into Database: "+err.Error())
			rw.Write([]byte("Can't insert data into Database"))
		}

	}else{
		rw.Write([]byte("not enough parameters given"))
	}
}
func get_all_function(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func get_func(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if len(url) != 0 {
		rows, err := db.Query("SELECT * FROM website WHERE url like ?", url)
		if (err != nil) {
			fmt.Println("can't execute select query: "+err.Error())
		}
		for rows.Next() {
			w := new(Website)
			err := rows.Scan(&w.ID, &w.Url, &w.UserName, &w.Password)
			if (err != nil) {
				fmt.Println("can't read into struct: "+err.Error())
			}
			rw.Write([]byte(w.ToString()))

		}
		err = rows.Err()
		if(err != nil){
			fmt.Println("Error with Row: "+err.Error())
		}
	}else{
		rw.Write([]byte("no get parameter given"))
	}
}
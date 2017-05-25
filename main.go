package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/context"
)

var db *sql.DB
var configuration Configuration
var key []byte

type Configuration struct {
	DBHost            string `json:"db_host"`
	DBUser            string `json:"db_user"`
	DBPassword        string `json:"db_password"`
	BasicAuthUser     string `json:"basic_auth_user"`
	BasicAuthPassword string `json:"basic_auth_password"`
	Database          string `json:"database"`
	Key               string `json:"key"`
}

func init() {
	fmt.Println("<Konfiguration lesen>")
	configFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Konfigurations Lesefehler: " + err.Error())
	}
	jsonDecoder := json.NewDecoder(configFile)
	configuration = Configuration{}
	jsonDecoder.Decode(&configuration)
	fmt.Println("<Datenbankverbindung herstellen>")
	db, err = sql.Open("mysql", configuration.DBUser+":"+configuration.DBPassword+"@tcp("+configuration.DBHost+")/"+configuration.Database+"?parseTime=true")
	if err != nil {
		fmt.Println("Datenbankzugriffs fehler: " + err.Error())
	}
	key = []byte(configuration.Key) // 32 bytes
}

func main() {
	router := NewRouter()
	http.Handle("/", router)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}

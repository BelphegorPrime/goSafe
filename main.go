package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"os"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/context"
)

var db *sql.DB
var configuration Configuration
var key []byte
type Configuration struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Key 	 string `json:"key"`
}

func init() {
	fmt.Println("<Konfiguration lesen>")
	configFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Konfigurations Lesefehler: "+err.Error())
	}
	jsonDecoder := json.NewDecoder(configFile)
	configuration = Configuration{}
	jsonDecoder.Decode(&configuration)
	fmt.Println("<Datenbankverbindung herstellen>")
	dbFromConfig, err := sql.Open("mysql", configuration.User+":"+configuration.Password+"@tcp("+configuration.Host+")/"+configuration.Database+"?parseTime=true")
	if err != nil {
		fmt.Println("Datenbankzugriffs fehler: "+err.Error())
	}
	db = dbFromConfig

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
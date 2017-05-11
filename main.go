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
type Configuration struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
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
}

func main() {
	router := NewRouter()
	http.Handle("/", router)
	errHTTP := http.ListenAndServe(":8080", nil)
	if errHTTP != nil {
		fmt.Println("Error: " + errHTTP.Error())
	}
}

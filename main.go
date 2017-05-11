package main

import (
	"fmt"
	"net/http"
	"runtime"
	"database/sql"
	"os"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
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
	configuration := Configuration{}
	jsonDecoder.Decode(&configuration)
	fmt.Println("<Datenbankverbindung herstellen>")
	dbFromConfig, err := sql.Open("mysql", configuration.User+":"+configuration.Password+"@tcp("+configuration.Host+")/"+configuration.Database+"?parseTime=true")
	if err != nil {
		fmt.Println("Datenbankzugriffs fehler: "+err.Error())
	}
	db = dbFromConfig
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

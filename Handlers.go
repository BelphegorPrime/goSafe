package main

import (
	"encoding/json"
	"fmt"
	"github.com/BelphegorPrime/lib"
	"net/http"
)

func index_func(rw http.ResponseWriter, req *http.Request){
	rw.Write([]byte("HelloWorld"))
}

func save_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := lib.GetRequestContentFromRequest(req)
	returnValue := Save(
		requestContent["url"].(string),
		requestContent["username"].(string),
		requestContent["password"].(string),
	)
	cipherText, err := lib.Encrypt(returnValue, key)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	values := map[string]interface{}{"responseText": string(cipherText)}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

func get_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := lib.GetRequestContentFromRequest(req)
	returnValue := []string{}
	if requestContent["crypto"].(float64) >= 0 {
		returnValue = Get(requestContent["url"].(string), requestContent["crypto"].(float64))
	}else{
		returnValue = Get(requestContent["url"].(string), -1)
	}

	values := map[string]interface{}{"responseText": returnValue}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

func delete_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := lib.GetRequestContentFromRequest(req)
	returnValue := Delete(
		requestContent["url"].(string),
		requestContent["username"].(string),
		requestContent["password"].(string),
	)
	cipherText, err := lib.Encrypt(returnValue, key)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	values := map[string]interface{}{"responseText": string(cipherText)}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

func all_func(rw http.ResponseWriter, req *http.Request) {
	returnValue := All()
	for i := 0; i < len(returnValue); i++ {
		cipherText, err := lib.Encrypt([]byte(returnValue[i]), key)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		returnValue[i] = string(cipherText)
	}
	values := map[string]interface{}{"responseText": returnValue}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

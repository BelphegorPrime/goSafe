package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func save_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	returnValue := Save(
		requestContent["url"].(string),
		requestContent["username"].(string),
		requestContent["password"].(string),
	)
	cipherText, err := encrypt(returnValue)
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
	requestContent := getRequestContentFromRequest(req)
	returnValue := Get(requestContent["url"].(string))
	for i := 0; i < len(returnValue); i++ {
		cipherText, err := encrypt([]byte(returnValue[i]))
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		returnValue[i]=string(cipherText)
	}
	values := map[string]interface{}{"responseText": returnValue}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

func delete_func(rw http.ResponseWriter, req *http.Request){
	requestContent := getRequestContentFromRequest(req)
	returnValue := Delete(
		requestContent["url"].(string),
		requestContent["username"].(string),
		requestContent["password"].(string),
	)
	cipherText, err := encrypt(returnValue)
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

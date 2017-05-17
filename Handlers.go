package main

import (
	"fmt"
	"net/http"
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
	rw.Write(cipherText)
}

func get_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	returnValue := Get(requestContent["url"].(string))
	cipherText, err := encrypt(returnValue)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(cipherText)
}

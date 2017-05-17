package main

import (
	"fmt"
	"net/http"
)

func save_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	returnValue := Save(requestContent)
	cipherText, err := encrypt(returnValue)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(cipherText)
}

func get_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := getRequestContentFromRequest(req)
	returnValue := Get(requestContent)
	cipherText, err := encrypt(returnValue)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(cipherText)
}

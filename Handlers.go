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

	returnValue = makeCrypto(returnValue, requestContent["crypto"].(float64))

	values := map[string]interface{}{"responseText": string(returnValue)}
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

	returnValue = makeCrypto(returnValue, requestContent["crypto"].(float64))

	values := map[string]interface{}{"responseText": string(returnValue)}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

func all_func(rw http.ResponseWriter, req *http.Request) {
	requestContent := lib.GetRequestContentFromRequest(req)
	returnValue := All(requestContent["crypto"].(float64))

	values := map[string]interface{}{"responseText": returnValue}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	rw.Write(jsonValue)
}

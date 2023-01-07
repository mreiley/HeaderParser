/*
Request Header Parser Microservice

AUTHOR: Mario Reiley
NOTE: it is my port to Golang

Build a full stack JavaScript app that is functionally similar to this:
https://request-header-parser-microservice.freecodecamp.rocks/

Tests
	* You should provide your own project, not the example URL.
	* A request to /api/whoami should return a JSON object with your IP address in the ipaddress key.
	* A request to /api/whoami should return a JSON object with your preferred language in the language key.
	* A request to /api/whoami should return a JSON object with your software in the software key.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Project json format
type jformat struct {
	Ipaddress string `json:"ipaddress"`
	Language  string `json:"language"`
	Software  string `json:"software"`
}

// just begin the services
func HandlerSendFile(w http.ResponseWriter, req *http.Request) {
	fileBytes, err := os.ReadFile("./views/index.html")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

// get data from de client , sometime ip,lang and soft do not come from the client.
func HandlerHeaderData(w http.ResponseWriter, req *http.Request) {
	var data jformat
	data.Ipaddress = (*req).Header.Get("X-Forwarded-For")
	data.Language = (*req).Header.Get("Accept-Language")
	data.Software = (*req).Header.Get("user-agent")
	res, _ := json.Marshal(data)
	w.Write(res)
}

func main() {
	http.HandleFunc("/", HandlerSendFile)
	http.HandleFunc("/api/whoami", HandlerHeaderData)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

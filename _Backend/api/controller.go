package main

import (
	"./services"
	"encoding/json"
	"fmt"
	"net/http"
)


func main() {
	PORT := ":4444"
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(PORT,nil)
}

func testHandler(write http.ResponseWriter, request *http.Request){
	//SHOULD USE CONFIG FILE OR ENV
	host := ""
	//a := services.GetIbopInfo(host)
	//fmt.Printf("Found a single document: %+v\n", a)
	a := services.GetRealTimeInfo(host)
	fmt.Printf("Found a single document: %+v\n", a)

	// PARSING TO JSON
	b, _ := json.Marshal(a)
	fmt.Fprint(write, string(b))
}
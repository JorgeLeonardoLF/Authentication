package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckServerHealth(w http.ResponseWriter, r *http.Request) {
	payload := struct{}{}
	data, err := json.Marshal(payload)
	if err != nil {
		/*
			If something goes wrong when trying to marshal the payload will will write to the response a header with a status code
			In this case 500 - Internal server error
			Log the issue
			return
		*/
		w.WriteHeader(500)
		log.Printf("Failed to marshal JSON response: %v", payload) // Note we are not fatal
		return
	}
	/*
		If no error on marshal
		Add a header to the response to say we are responding with json data
		"Content-Type", "application/json" --> Key/ value pair this is the standard for json responses
		TLDR; this lets it know what to expect
	*/
	w.Header().Add("Content-Type", "application/json")

	/*Now we write status code that was passed to us*/
	w.WriteHeader(200)
	/*Write the data it self*/
	w.Write(data)
}

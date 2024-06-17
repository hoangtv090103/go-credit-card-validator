package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Valid bool `json:"valid"`
}

func CreateCardValidator(writer http.ResponseWriter, request *http.Request) {
	var cardNumber struct {
		Number string `json:"number"` // Number field holds the credit card number.
	}

	err := json.NewDecoder(request.Body).Decode(&cardNumber)

	fmt.Println("Number: ", cardNumber.Number)

	if err != nil {
		http.Error(writer, "Invalid JSON payload", http.StatusBadRequest)
	}

	// Validate the credit card number using the Luhn algorithm.
	isValid := Validator(cardNumber.Number)

	// Create a response struct with the validation result.
	response := Response{
		Valid: isValid,
	}

	// Marshal the response struct into JSON format.
	jsonResponse, err := json.Marshal(response) // Marshal:
	if err != nil {
		http.Error(writer, "Error creating response", http.StatusInternalServerError)
		return
	}

	// Set the content type header to indicate JSON response.
	writer.Header().Set("Content-Type", "application/json")

	writer.Write(jsonResponse)

}

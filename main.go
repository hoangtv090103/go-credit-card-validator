package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/a-h/templ"
	"github.com/joho/godotenv"
	// "credit-card-validator/views"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	// component := views.Home()
	// http.Handle("/", templ.Handler(component))
	http.HandleFunc("/", CreateCardValidator)
	fmt.Println("Listening on port: ", port)

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
		
}
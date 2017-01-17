package main

import (
	"log"
	"os"

	"github.com/JKhawaja/sendinblue"
)

func main() {
	// recommendation: set API key as system environment variable
	apiKey := os.Getenv("SIB_KEY")

	// Create SendInBlue Client
	sibClient, err := sib.NewClient(apiKey)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Create Email
	email := sib.NewEmail()
	email.From[0] = "user@example.net" // ADD SENDER EMAIL HERE
	email.From[1] = "User"
	email.Subject = "Test"
	email.To["user@example.net"] = "User Name" // CHANGE TO TEST DELIVERY ADDRESS
	email.Text = "Hello World."

	// Send Email
	response, err := sibClient.SendEmail(email)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Print Response
	log.Println(response.Code)
	log.Println(response.Message)
}

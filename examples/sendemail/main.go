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
	email.From = [2]string{"sender@example.net", "Sender Name"}
	email.Subject = "Test"
	email.To["user1@example.net"] = "User 1"
	// email.To["user2@example.net"] = "User 2"
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

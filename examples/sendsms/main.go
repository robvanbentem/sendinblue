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

	// Create SMS
	request := &sib.SMSRequest{
		To:   "", // Mobile Number to send to
		From: "Sender",
		Text: "Hello World.",
		Type: "marketing",
	}

	// Send SMS
	resp, err := sibClient.SendSMS(request)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Print Response
	log.Println(resp.Code)
	log.Println(resp.Message)
}

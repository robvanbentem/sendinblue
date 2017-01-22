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
		log.Println("Client creation error: ")
		log.Println(err)
		os.Exit(1)
	}

	/* Create SMS Campaign */
	myCampaign := &sib.SMSCampaign{
		Name:     "Test SMS Campaign",
		Sender:   "Tester",
		Content:  "Hello World",
		Send_now: 1, // ready to send
	}

	createResp, err := sibClient.CreateSMSCampaign(myCampaign)
	if err != nil {
		log.Println("Create campaign error: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println(createResp.Code)
	log.Println(createResp.Message)

	/* Update SMS Campaign */
	updateCampaign := &sib.SMSCampaign{
		Name:     "Test SMS Campaign",
		Sender:   "Tester",
		Content:  "Hello World. UPDATED!",
		Send_now: 1,
	}

	err = sibClient.UpdateSMSCampaign(createResp.Data.Id, updateCampaign)
	if err != nil {
		log.Println("Update campaign error: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Update Campaign: successful")

	/* Test SMS Campaign */
	testNumber := "+00000000000" // Mobile Number to test campaign on
	testResp, err := sibClient.SMSCampaignTest(createResp.Data.Id, testNumber)
	if err != nil {
		log.Println("Error while testing campaign: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println(testResp.Code)
	log.Println(testResp.Message)
}

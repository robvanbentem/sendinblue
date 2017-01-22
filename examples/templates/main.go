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

	/* Create Template */
	myTemplate := &sib.Template{
		Template_name: "Test Template",
		Html_content:  "Hello World.",
		Subject:       "Test Template Email",
		From_email:    "sender@example.net", // SENDER EMAIL HERE
		Status:        1,                    // activate template
	}

	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println("Create template error: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println(createResponse.Code)
	log.Println(createResponse.Message)

	/* Send Template Email */
	templateID := createResponse.Data.ID

	userList := []string{"user1@example.net", "user2@example.net", "user3@example.net"} // RECIEVER EMAILS HERE

	sendResponse, err := sibClient.SendTemplateEmail(templateID, userList, nil)
	if err != nil {
		log.Println("Send template email error: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println(sendResponse.Code)
	log.Println(sendResponse.Message)

	/* Update Template */
	udpateTemplate := &sib.Template{
		Template_name: "Test Template",
		Html_content:  "Hello World. UPDATED!",
		Subject:       "Test Template Email",
		From_email:    "sender@example.net", // SENDER EMAIL HERE
		Status:        1,
	}

	err = sibClient.UpdateTemplate(templateID, udpateTemplate)
	if err != nil {
		log.Println("Update template error: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Update Template: succesful")

	/* Get Template */
	getResponse, err := sibClient.GetTemplate(templateID)
	if err != nil {
		log.Println("Get template error: ")
		log.Println(err)
		os.Exit(1)
	}

	log.Println(getResponse)
}

package sib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	apiKey string
	Client *http.Client
}

func NewClient(apiKey string) (*Client, error) {

	emptyClient := &Client{}

	if apiKey == "" {
		err := fmt.Errorf("Please provide a SendInBlue API Key.")
		return emptyClient, err
	}

	return &Client{
		apiKey: apiKey,
		Client: &http.Client{},
	}, nil
}

func (c *Client) SendEmail(e *Email) (EmailResponse, error) {

	emptyResp := EmailResponse{}

	body, err := json.Marshal(e)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v2.0/email", r)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response EmailResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Error: could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

func (c *Client) CreateTemplate(t *Template) (TemplateResponse, error) {

	emptyResp := TemplateResponse{}

	body, err := json.Marshal(t)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v2.0/template", r)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response TemplateResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Error: could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

func (c *Client) UpdateTemplate(id string, t *Template) (TemplateResponse, error) {

	emptyResp := TemplateResponse{}

	body, err := json.Marshal(t)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/template/%s", id)
	req, err := http.NewRequest("POST", url, r)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response TemplateResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Error: could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

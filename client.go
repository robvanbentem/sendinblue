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

func (c *Client) GetTemplate(template_id int) (CampaignResponse, error) {

	emptyResp := CampaignResponse{}

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/campaign/%s/detailsv2", template_id)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response CampaignResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Error: could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

// Takes a Type, Status, Page, and Page_Limit as arguments
func (c *Client) ListTemplates(t *TemplateList) (TemplateListResponse, error) {

	emptyResp := TemplateListResponse{}

	body, err := json.Marshal(t)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/campaign/detailsv2")
	req, err := http.NewRequest("GET", url, r)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response TemplateListResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Error: could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

func (c *Client) SendTemplateEmail(id int, to []string, e *EmailOptions) (EmailResponse, error) {

	toString := ""

	if len(to) > 1 {
		for _, val := range to {
			toString = toString + val + "|"
		}
	} else {
		toString = toString + to[0]
	}

	email := TemplateEmail{}
	email.To = toString

	if e != nil {
		email.Cc = e.Cc
		email.Bcc = e.Bcc
		email.ReplyTo = e.ReplyTo
		email.Attr = e.Attr
		email.Attachment_url = e.Attachment_url
		email.Attachment = e.Attachment
		email.Headers = e.Headers
	}

	emptyResp := EmailResponse{}

	body, err := json.Marshal(email)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/template/%s", id)
	req, err := http.NewRequest("PUT", url, r)
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

func (c *Client) UpdateTemplate(id int, t *Template) (TemplateResponse, error) {

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

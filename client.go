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

func (c *Client) AggregateReport(a *AggregateReport) (AggregateResponse, error) {

	emptyResp := AggregateResponse{}

	body, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v2.0/statistics", r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response AggregateResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

func (c *Client) CreateSMSCampaign(s *SMSCampaign) (SMSCampaignResponse, error) {

	emptyResp := SMSCampaignResponse{}

	body, err := json.Marshal(s)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v2.0/sms", r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response SMSCampaignResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Could not decode response format: ", err)
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
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
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
		err := fmt.Errorf("Could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

// Start and End dates must be in YYYY-MM-DD format
// Start date must be before end date, and end date must be after start date
func (c *Client) DeleteBouncedEmails(start, end, email string) error {

	request := DeleteBouncesRequest{
		Start_date: start,
		End_date:   end,
		Email:      email,
	}

	body, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return err
	}
	r := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v2.0/bounces", r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("Request error: ", resp.Status)
		return err
	}

	return nil
}

func (c *Client) GetTemplate(template_id int) (CampaignResponse, error) {

	emptyResp := CampaignResponse{}

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/campaign/%v/detailsv2", template_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
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
		err := fmt.Errorf("Could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

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
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
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
		err := fmt.Errorf("Could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
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
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
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
		err := fmt.Errorf("Could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

func (c *Client) SendSMS(s *SMSRequest) (SMSResponse, error) {

	emptyResp := SMSResponse{}

	body, err := json.Marshal(s)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return emptyResp, err
	}
	r := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v2.0/sms", r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("Could not recognize API response format: ", err)
		return emptyResp, err
	}

	var response SMSResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(resp.Body)
		err := fmt.Errorf("Could not decode response format: ", err)
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
		toString = to[0]
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

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/template/%v", id)
	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return emptyResp, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return emptyResp, err
	}
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
		err := fmt.Errorf("Could not decode response format: ", err)
		return emptyResp, err
	}

	return response, nil
}

func (c *Client) UpdateSMSCampaign(id int, s *SMSCampaign) error {

	body, err := json.Marshal(s)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return err
	}
	r := bytes.NewReader(body)

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/sms/%v", id)
	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("Request error: ", resp.Status)
		return err
	}

	return nil
}

func (c *Client) UpdateTemplate(id int, t *Template) error {

	body, err := json.Marshal(t)
	if err != nil {
		err = fmt.Errorf("Could not marshal JSON: ", err)
		return err
	}
	r := bytes.NewReader(body)

	url := fmt.Sprintf("https://api.sendinblue.com/v2.0/template/%v", id)
	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		err := fmt.Errorf("Could not create http request: ", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		err := fmt.Errorf("Could not send http request: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("Request error: ", resp.Status)
		return err
	}

	return nil
}

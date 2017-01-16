package sib

import (
	"bufio"
	"encoding/base64"
	"os"
)

// API Docs: https://apidocs.sendinblue.com/tutorial-sending-transactional-email/
type Email struct {
	To           map[string]string `json:"to"`
	Subject      string            `json:"subject"`
	From         [2]string         `json:"from"`
	HTML         string            `json:"html"`
	Text         string            `json:"text"`
	CC           map[string]string `json:"cc"`
	Bcc          map[string]string `json:"bcc"`
	ReplyTo      map[string]string `json:"replyto"`
	Attachment   map[string]string `json:"attachment"` // must be URL
	Headers      map[string]string `json:"headers"`
	Inline_image map[string]string `json:"inline_image"`
}

// API Docs: https://apidocs.sendinblue.com/template/
type Template struct {
	From_name      string
	Template_name  string // Mandatory
	Bat            string
	Html_content   string // Mandatory if no html_url
	Html_url       string // Mandatory if no html_content
	Subject        string
	From_email     string
	Reply_to       string
	To_field       string
	Status         int // 0 (active) or 1 (inactive)
	Attachment_url string
}

type TemplateList struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	Page       int    `json:"page"`
	Page_limit int    `json:"page_limit"`
}

type EmailData struct {
	Message_id string `json:"message-id"`
}

type EmailResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    EmailData `json:"data"`
}

type TemplateData struct {
	ID string `json:"id"`
}

type TemplateResponse struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    TemplateData `json:"data"`
}

type CampaignData struct {
	ID            string `json:"id"`
	Campaign_name string `json:"campaign_name"`
	Subject       string `json:"subject"`
	Bat_sent      string `json:"bat_sent"`
	Type          string `json:"type"`
	Html_content  string `json:"html_content"`
	Entered       string `json:"entered"`
	Modified      string `json:"modified"`
	Templ_status  string `json:"templ_status"`
	From_name     string `json:"from_name"`
	From_email    string `json:"from_email"`
	Reply_to      string `json:"reply_to"`
	To_field      string `json:"to_field"`
}

type CampaignResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    []CampaignData `json:"data"`
}

type TemplateListData struct {
	Campaign_records       []CampaignData `json:"campaign_records"`
	Page                   int            `json:"page"`
	Page_limit             int            `json:"page_limit"`
	Total_campaign_records int            `json:"total_campaign_records"`
}

type TemplateListResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    TemplateListData `json:"data"`
}

func NewEmail() *Email {

	var from [2]string
	attachment := make(map[string]string)
	bcc := make(map[string]string)
	cc := make(map[string]string)
	headers := make(map[string]string)
	to := make(map[string]string)
	inline_image := make(map[string]string)

	return &Email{
		Attachment:   attachment,
		Bcc:          bcc,
		CC:           cc,
		From:         from,
		Headers:      headers,
		To:           to,
		Inline_image: inline_image,
	}

	return nil
}

// AddImage() Method ...
// returns the filename, which can (and should) be used as a variable in HTML
// < img src="{{{filename}}}" alt="image" border="0" >
func (e *Email) AddImage(f *os.File) string {
	filename := f.Name()

	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	fileReader := bufio.NewReader(f)
	fileReader.Read(buffer)

	e.Inline_image[filename] = base64.StdEncoding.EncodeToString(buffer)

	return filename
}

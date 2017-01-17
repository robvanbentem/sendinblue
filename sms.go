package sib

/* Request Types */

type SMSRequest struct {
	To      string `json:"to"`   // Mobile Number (Mandatory)
	From    string `json:"from"` // No more than 11 alphanumeric characters (Mandatory)
	Text    string `json:"text"` // No more than 160 characters (Mandatory)
	Web_url string `json:"web_url"`
	Tag     string `json:"tag"`
	Type    string `json:"type"` // "marketing" (default) or "transactional"
}

/* Response Types */

type SMSReference struct {
	One string `json:"1"`
}

type SMSData struct {
	Status           string       `json:"status"`
	Number_sent      int          `json:"number_sent"`
	To               string       `json:"to"`
	Sms_count        int          `json:"sms_count"`
	Credits_used     float64      `json:"json:"credits_used"`
	Remaining_credit float64      `json:"remaining_credit"`
	Reference        SMSReference `json:"reference"`
	Description      string       `json:"description"`
	Reply            string       `json:"reply"`
	Bounce_type      string       `json:"bounce_type"`
	Error_code       int          `json:"error_code"`
}

type SMSResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    SMSData `json:"data"`
}

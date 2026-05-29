package email

type SenderDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReceiverDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailRequest struct {
	Sender      SenderDetails   `json:"sender"`
	Receiver    ReceiverDetails `json:"receiver"`
	Subject     string          `json:"subject"`
	HTMLContent string          `json:"html_content"`
}

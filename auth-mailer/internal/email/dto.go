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
	Sender      SenderDetails
	Receiver    ReceiverDetails
	Subject     string
	HTMLContent string
}

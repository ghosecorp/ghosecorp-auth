package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	apiKey   string
	emailURL string
	client   *http.Client
}

func NewService(apiKey string, emailURL string) *Service {
	return &Service{
		apiKey:   apiKey,
		emailURL: emailURL,
		client:   &http.Client{},
	}
}

func (s *Service) SendEmail(
	ctx context.Context,
	req EmailRequest,
) (map[string]interface{}, error) {

	payload := map[string]interface{}{
		"sender": map[string]string{
			"name":  req.Sender.Name,
			"email": req.Sender.Email,
		},
		"to": []map[string]string{
			{
				"name":  req.Receiver.Name,
				"email": req.Receiver.Email,
			},
		},
		"subject":     req.Subject,
		"htmlContent": req.HTMLContent,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		s.emailURL,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Set("accept", "application/json")
	request.Header.Set("api-key", s.apiKey)
	request.Header.Set("content-type", "application/json")

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&result)

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("brevo api error")
	}

	return result, nil
}

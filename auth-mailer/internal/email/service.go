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
) (map[string]any, error) {

	payload := map[string]any{
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

	var result map[string]any
	decodeErr := json.NewDecoder(response.Body).Decode(&result)

	if response.StatusCode >= 400 {
		if decodeErr == nil {
			if msg, ok := result["message"].(string); ok {
				return nil, fmt.Errorf("brevo api error (status %d): %s", response.StatusCode, msg)
			}
			return nil, fmt.Errorf("brevo api error (status %d): %v", response.StatusCode, result)
		}
		return nil, fmt.Errorf("brevo api error with status %d", response.StatusCode)
	}

	if decodeErr != nil {
		return nil, decodeErr
	}

	return result, nil
}

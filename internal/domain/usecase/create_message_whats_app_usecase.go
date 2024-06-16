package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/reangeline/micro_saas/internal/dto"
)

type CreateMessageWhatsAppUseCase struct {
}

func NewCreateMessageWhatsAppUseCase() *CreateMessageWhatsAppUseCase {
	return &CreateMessageWhatsAppUseCase{}
}

func (cwa *CreateMessageWhatsAppUseCase) Execute(ctx context.Context, input *dto.MessagePayload) error {

	if len(input.Entry) > 0 && len(input.Entry[0].Changes) > 0 && len(input.Entry[0].Changes[0].Value.Messages) > 0 {
		phoneID := input.Entry[0].Changes[0].Value.Metadata.PhoneNumberID
		from := input.Entry[0].Changes[0].Value.Messages[0].From
		body := input.Entry[0].Changes[0].Value.Messages[0].Text.Body

		acknowledgeMessage(phoneID, from, body)
	}

	return nil
}

// AcknowledgeMessage sends an acknowledgement message
func acknowledgeMessage(phoneNumberID, from, body string) {

	url := ""

	reply := TesteOpenai(body)

	var msg dto.MessageCompletion

	// Deserializando a string JSON para a struct MessageCompletion
	err := json.Unmarshal([]byte(reply), &msg)
	if err != nil {
		fmt.Printf("Erro ao deserializar JSON: %s\n", err)
		return
	}

	// Construct the JSON payload
	messageData := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                "5511967700232",
		"text": map[string]interface{}{
			"body": msg.Choices[0].Message.Content,
		},
	}

	// Convert messageData to JSON
	jsonData, err := json.Marshal(messageData)
	if err != nil {
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	// Set the required headers
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	log.Println(body)

}

func TesteOpenai(replyOpenAi string) string {
	apiKey := ""

	url := "https://api.openai.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a normal person.",
			},
			{
				"role":    "user",
				"content": replyOpenAi,
			},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error occurred during request creation. Error: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error occurred during request execution. Error: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occurred during reading the response body. Error: %s", err.Error())
	}

	return string(body)

}

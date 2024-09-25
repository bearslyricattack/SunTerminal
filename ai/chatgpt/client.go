package chatgpt

import (
	"Aterminal/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func Query(path string, input string) (string, error) {
	model, err := config.GetModelInfo(path)
	if err != nil {
		return "", fmt.Errorf("get model info error: %w", err)
	}
	apiKey := model.Key
	chatRequest := ChatRequest{
		//Model: "gpt-3.5-turbo",
		Model: model.Type,
		Messages: []Message{
			{Role: "user", Content: input},
		},
	}
	reqBody, err := json.Marshal(chatRequest)
	if err != nil {
		fmt.Println("marshal error:", err)
		return "", fmt.Errorf("marshal error: %w", err)
	}
	req, err := http.NewRequest("POST", model.Path, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("create request error:", err)
		return "", fmt.Errorf("create request error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request error:", err)
		return "", fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read result error:", err)
		return "", fmt.Errorf("read result error: %w", err)
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return "", fmt.Errorf("unmarshal error: %w", err)
	}
	if len(chatResponse.Choices) > 0 {
		return chatResponse.Choices[0].Message.Content, nil
	} else {
		return "", fmt.Errorf("respense error: %w", err)
	}
}

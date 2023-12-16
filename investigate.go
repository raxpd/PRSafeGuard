package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func investigation(input string) bool {
	if input == "skip" {
		fmt.Println("Skipping OpenAI call.")
		return true
	}

	apiKeyBytes, err := ioutil.ReadFile("openaikey.pem")
	if err != nil {
		fmt.Println("Error reading API key:", err)
	}
	apiKey := strings.TrimSpace(string(apiKeyBytes))

	body := RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: input,
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshalling body:", err)
		return false
	}

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false
	}

	type Response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	var respContent Response
	err = json.Unmarshal(responseBody, &respContent)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return false
	}

	scoreStr := respContent.Choices[0].Message.Content
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		fmt.Println("Error converting score to integer:", err)
		return false
	}

	if score > 50 {
		return false
	}

	return true
}

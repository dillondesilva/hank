package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slfy/internal/utils"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func CallLMStudioLLM(messages []Message) (string, error) {
	url := "http://127.0.0.1:1234/v1/chat/completions"

	requestBody := ChatRequest{
		Model:    "lmstudio-llm",
		Messages: messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResp ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		return "", err
	}

	return chatResp.Choices[0].Message.Content, nil
}

func BuildPromptFromCommandAndError(command string, errOut string, stdOut string, threshold ...int) string {
	// Set default threshold
	thresh := 4000
	if len(threshold) > 0 {
		thresh = threshold[0]
	}

	// Use thresh instead of hardcoded 4000
	if len(stdOut) > thresh {
		stdOut = stdOut[:thresh]
	}

	userCommand := fmt.Sprintf("<command>%s</command>", command)
	userStdOut := fmt.Sprintf("<stdout>Here are the last %d characters of the standard output:\n%s</stdout>", thresh, stdOut)
	userError := fmt.Sprintf("<error>%s</error>", errOut)
	combinedPrompt := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s", utils.DefaultAssistancePrompt, userCommand, userStdOut, userError)
	return combinedPrompt
}

package ai_service

import (
	"blogv2/global"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/siddontang/go/log"
	"io"
	"net/http"
	"strings"
)

// 定义请求数据结构
type ChatCompletionRequest struct {
	Model       string                  `json:"model"`
	Messages    []ChatCompletionMessage `json:"messages"`
	MaxTokens   int                     `json:"max_tokens,omitempty"`
	Temperature float64                 `json:"temperature,omitempty"`
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 定义响应数据结构
type ChatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func AiChat(context string) (resContent string, err error) {
	apiKey := global.Config.Ai.SecretKey
	//url := "https://api.deepseek.com/v1/chat/completions"
	url := "https://api.chatanywhere.tech/v1/chat/completions"
	// 构造请求体
	requestBody := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatCompletionMessage{
			{
				Role:    "user",
				Content: "我只会给你一篇文章的正文，你需要从中分析出与正文相关的标题，简介，分类和标签，并且按照固定格式进行输出为文本，\n不管我问什么，你都要返回如下格式 {\"title\": \"\", \"abstract\": \"\", \"category\": \"\", \"tag\": []}",
			},
			{
				Role:    "user",
				Content: context,
			},
		},
		Temperature: 0.1,
	}

	// 序列化为JSON
	jsonBody, _ := json.Marshal(requestBody)

	// 创建HTTP请求
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, _ := io.ReadAll(resp.Body)
	var response ChatCompletionResponse
	json.Unmarshal(body, &response)

	// 处理结果
	if resp.StatusCode == http.StatusOK {
		if len(response.Choices) > 0 {
			resContent = response.Choices[0].Message.Content
		} else {
			err = errors.New("未返回有效结果")
			return
		}
	} else {
		log.Error("请求失败，状态码: %d\n错误信息: %s\n", resp.StatusCode, response.Error.Message)
		err = errors.New(fmt.Sprintf("请求失败，状态码: %d\n错误信息: %s\n", resp.StatusCode, response.Error.Message))
		return
	}
	return
}

package ai_service

import (
	"blogv2/common/res"
	"blogv2/global"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type Choice struct {
	Index int `json:"index"`
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason interface{} `json:"finish_reason"`
}

type StreamData struct {
	Id                string      `json:"id"`
	Choices           []Choice    `json:"choices"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	Object            string      `json:"object"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

func ChatStream(list []llms.MessageContent, c *gin.Context) (err error) {
	llm, err := openai.New(
		openai.WithModel("gpt-3.5-turbo"),
		openai.WithBaseURL("https://api.chatanywhere.tech"),
		openai.WithToken(global.Config.Ai.SecretKey),
	)

	if err != nil {
		logrus.Error(err)
	}
	messages := list
	_, err = llm.GenerateContent(context.Background(), messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		res.SSESuccess(c, string(chunk))
		return nil
	}))
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

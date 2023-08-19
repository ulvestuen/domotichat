package chat

import (
	"context"
	"domotichat-api/internal/config"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var systemMessage = []openai.ChatCompletionMessage{
	{
		Role: openai.ChatMessageRoleSystem,
		Content: "Please answer in an enthusiastic tone, with great variety in language artifacts." +
			"If there are elements you are not exactly sure about or you find the question ambiguous, you shall ask back to gather more " +
			"information.",
	},
}

func SendChatMessage(messages []openai.ChatCompletionMessage) (messageHistory []openai.ChatCompletionMessage) {
	messageHistory = messages
	openAiConfig := openai.DefaultConfig(config.OpenAiApiKey)
	openAiConfig.OrgID = config.OpenAiOrg
	client := openai.NewClientWithConfig(openAiConfig)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT4,
			Messages:  append(systemMessage, messageHistory...),
			Functions: functionDefinitions,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return messages
	}
	assistantMessage := resp.Choices[0].Message
	messageHistory = append(messageHistory, assistantMessage)

	if functionCall := assistantMessage.FunctionCall; functionCall != nil {
		functionsMessage := functionMappings[functionCall.Name](nil)
		messageHistory = append(messageHistory, functionsMessage)
		var functionResponse, err = client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:     openai.GPT4,
				Messages:  append(systemMessage, messageHistory...),
				Functions: functionDefinitions,
			},
		)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return messages
		}
		return append(messageHistory, functionResponse.Choices[0].Message)
	}
	return messageHistory
}

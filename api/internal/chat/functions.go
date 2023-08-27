package chat

import (
	"domotichat-api/internal/domotics"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
)

var functionDefinitions = []openai.FunctionDefinition{
	{
		Name:        "turn_light_bulb_on",
		Description: "Turn on a light bulb in my smart home system",
		Parameters:  json.RawMessage(`{"type":"object","properties":{}}`),
	},
	{
		Name:        "turn_light_bulb_off",
		Description: "Turn off a light bulb in my smart home system",
		Parameters:  json.RawMessage(`{"type":"object","properties":{}}`),
	},
}

var functionMappings = map[string]func(any) openai.ChatCompletionMessage{
	"turn_light_bulb_on":  turnLightBulbOn,
	"turn_light_bulb_off": turnLightBulbOff,
}

func turnLightBulbOn(any) openai.ChatCompletionMessage {
	content, _ := domotics.ControlLight("garden_light_on")

	return openai.ChatCompletionMessage{
		Role:    "function",
		Name:    "turn_light_bulb_on",
		Content: content,
	}
}

func turnLightBulbOff(any) openai.ChatCompletionMessage {
	content, _ := domotics.ControlLight("garden_light_off")

	return openai.ChatCompletionMessage{
		Role:    "function",
		Name:    "turn_light_bulb_off",
		Content: content,
	}
}

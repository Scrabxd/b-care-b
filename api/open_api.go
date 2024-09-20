package api

import (
	helpers "b-care/Helpers"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func ConnectApi() *openai.Client {
	api_key := helpers.GetEnv("API_KEY")
	client := openai.NewClient(option.WithAPIKey(api_key))

	return client
}

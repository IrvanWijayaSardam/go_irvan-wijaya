package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type Input struct {
	Budget  int    `json:"budget" form:"budget"`
	Purpose string `json:"purpose" form:"purpose"`
	Brand   string `json:"brand" form:"brand"`
	RAM     string `json:"ram" form:"ram"`
	CPU     string `json:"cpu" form:"cpu"`
	Display string `json:"display" form:"display"`
}

var endpoint = "https://api.openai.com/v1/chat/completions"

func main() {
	godotenv.Load(".env")

	key := getAPIKey()

	e := echo.New()

	e.POST("/recommendation", func(c echo.Context) error {
		result, err := getResponseFromVA(c, key)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, response("failed", "Something Went Wrong!", err))
		}

		answer := openai.ChatCompletionMessage{
			Content: result.Choices[0].Message.Content,
		}

		return c.JSON(http.StatusOK, response("success", "", answer.Content))
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func getResponseFromVA(c echo.Context, key string) (openai.ChatCompletionResponse, error) {
	var input Input

	c.Bind(&input)

	var content = fmt.Sprintf("Berikan saya rekomendasi laptop dengan spesifikasi ini '''harga dalam rupiah: %d, tujuan: %s, merk: %s, RAM: %s, CPU: %s, layar: %s'''",
		input.Budget, input.Purpose, input.Brand, input.RAM, input.CPU, input.Display)

	client := openai.NewClient(key)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	return resp, err
}

func getAPIKey() string {
	var api_key string
	if result, find := os.LookupEnv("API_KEY"); find {
		api_key = result
	}

	if api_key == "" {
		log.Fatal("No API Key Provided.")
	}

	return api_key
}

func response(status, message string, data any) map[string]any {
	resp := map[string]any{
		"status": status,
	}

	if message != "" {
		resp["message"] = message
	}

	if data != nil {
		resp["data"] = data
	}

	return resp
}
